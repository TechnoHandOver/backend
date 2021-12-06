\c handover;

CREATE TYPE DAY_OF_WEEK AS ENUM ('Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'); --TODO: может всё-таки есть какой-то встроенный тип?

CREATE TABLE user_ (
    id SERIAL PRIMARY KEY,
    vk_id INT NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL CHECK (length(name) >= 2),
    avatar VARCHAR(500) NOT NULL
);

CREATE TABLE ad (
    id SERIAL PRIMARY KEY,
    user_author_id INT NOT NULL REFERENCES user_ (id) ON DELETE CASCADE,
    user_author_vk_id INT NOT NULL,
    user_author_name VARCHAR(100) NOT NULL,
    user_author_avatar VARCHAR(500) NOT NULL,
    user_executor_vk_id INT DEFAULT NULL,
    loc_dep VARCHAR(100) NOT NULL CHECK (length(loc_dep) >= 2),
    loc_arr VARCHAR(100) NOT NULL CHECK (length(loc_arr) >= 2),
    date_time_arr TIMESTAMP NOT NULL,
    item VARCHAR(50) NOT NULL CHECK (length(item) >= 3),
    min_price INT NOT NULL CHECK (min_price >= 0),
    comment VARCHAR(100) NOT NULL
);

CREATE TABLE ad_user_execution (
    ad_id INT NOT NULL PRIMARY KEY REFERENCES ad (id) ON DELETE CASCADE,
    user_executor_id INT NOT NULL REFERENCES user_ (id) ON DELETE CASCADE
);

CREATE TABLE route (
    id SERIAL PRIMARY KEY,
    user_author_id INT NOT NULL REFERENCES user_ (id) ON DELETE CASCADE,
    loc_dep VARCHAR(100) NOT NULL,
    loc_arr VARCHAR(100) NOT NULL,
    min_price INT NOT NULL CHECK (min_price >= 0)
);

CREATE TABLE route_tmp (
    id INT NOT NULL PRIMARY KEY REFERENCES route (id) ON DELETE CASCADE,
    date_time_dep TIMESTAMP NOT NULL,
    date_time_arr TIMESTAMP NOT NULL
);

CREATE TABLE route_perm (
    id INT NOT NULL PRIMARY KEY REFERENCES route (id) ON DELETE CASCADE,
    even_week BOOLEAN NOT NULL,
    odd_week BOOLEAN NOT NULL,
    day_of_week DAY_OF_WEEK NOT NULL,
    time_dep TIMESTAMP NOT NULL,
    time_arr TIMESTAMP NOT NULL
);

CREATE VIEW view_route_tmp (id, user_author_id, loc_dep, loc_arr, min_price, date_time_dep, date_time_arr)
    AS SELECT route.id, route.user_author_id, route.loc_dep, route.loc_arr, route.min_price, route_tmp.date_time_dep,
              route_tmp.date_time_arr
    FROM route
        JOIN route_tmp ON route.id = route_tmp.id
    ORDER BY route_tmp.date_time_dep, route_tmp.date_time_arr, route.min_price DESC, route.id;

CREATE VIEW view_route_perm (id, user_author_id, loc_dep, loc_arr, min_price, even_week, odd_week, day_of_week,
                             time_dep, time_arr)
    AS SELECT route.id, route.user_author_id, route.loc_dep, route.loc_arr, route.min_price, route_perm.even_week,
              route_perm.odd_week, route_perm.day_of_week, route_perm.time_dep, route_perm.time_arr
    FROM route
        JOIN route_perm ON route.id = route_perm.id
    ORDER BY route_perm.day_of_week, route_perm.time_dep, route_perm.time_arr, route.min_price DESC,
             route_perm.odd_week DESC, route_perm.even_week DESC, route.id;

CREATE FUNCTION user__update()
    RETURNS TRIGGER
AS $$
BEGIN
    UPDATE ad SET user_author_name = new.name, user_author_avatar = new.avatar
    WHERE user_author_id = new.id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER user__update AFTER UPDATE
    ON user_
    FOR EACH ROW
EXECUTE FUNCTION user__update();

CREATE FUNCTION ad_insert()
    RETURNS TRIGGER
AS $$
DECLARE user__ user_%ROWTYPE;
BEGIN
    SELECT INTO user__ id, vk_id, name, avatar FROM user_ WHERE id = new.user_author_id; --TODO: возможны ли оптимизации?
    new.user_author_vk_id = user__.vk_id;--(SELECT user_.vk_id FROM user_ WHERE user_.id = new.user_author_id);
    new.user_author_name = user__.name;
    new.user_author_avatar = user__.avatar;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ad_insert BEFORE INSERT
    ON ad
    FOR EACH ROW
EXECUTE FUNCTION ad_insert();

CREATE FUNCTION ad_user_execution_insert()
    RETURNS TRIGGER
AS $$
BEGIN
    UPDATE ad SET user_executor_vk_id = (SELECT user_.vk_id FROM user_ WHERE user_.id = new.user_executor_id)
    WHERE id = new.ad_id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ad_user_execution_insert AFTER INSERT
    ON ad_user_execution
    FOR EACH ROW
EXECUTE FUNCTION ad_user_execution_insert();

CREATE FUNCTION ad_user_execution_delete()
    RETURNS TRIGGER
AS $$
BEGIN
    UPDATE ad SET user_executor_vk_id = NULL
    WHERE id = old.ad_id;
    RETURN old;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ad_user_execution_delete AFTER DELETE
    ON ad_user_execution
    FOR EACH ROW
EXECUTE FUNCTION ad_user_execution_delete();

CREATE FUNCTION view_route_tmp_insert()
    RETURNS TRIGGER
AS $$
DECLARE id_ route.id%TYPE;
BEGIN
    INSERT INTO route (user_author_id, loc_dep, loc_arr, min_price)
    VALUES (new.user_author_id, new.loc_dep, new.loc_arr, new.min_price)
    RETURNING id INTO id_;
    INSERT INTO route_tmp (id, date_time_dep, date_time_arr)
    SELECT id_, new.date_time_dep, new.date_time_arr;
    new.id := id_;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_tmp_insert INSTEAD OF INSERT
    ON view_route_tmp
    FOR EACH ROW
    EXECUTE FUNCTION view_route_tmp_insert();

CREATE FUNCTION view_route_tmp_update()
    RETURNS TRIGGER
AS $$
BEGIN
    IF old.user_author_id != new.user_author_id THEN
        RAISE 'It is forbidden to update author of temporary route';
    END IF;
    UPDATE route SET loc_dep = new.loc_dep, loc_arr = new.loc_arr, min_price = new.min_price
    WHERE id = new.id AND user_author_id = new.user_author_id;
    UPDATE route_tmp SET date_time_dep = new.date_time_dep, date_time_arr = new.date_time_arr
    WHERE id = new.id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_tmp_update INSTEAD OF UPDATE
    ON view_route_tmp
    FOR EACH ROW
EXECUTE FUNCTION view_route_tmp_update();

CREATE FUNCTION view_route_tmp_delete()
    RETURNS TRIGGER
AS $$
BEGIN
    DELETE FROM route WHERE id = old.id;
    RETURN old;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_tmp_delete INSTEAD OF DELETE
    ON view_route_tmp
    FOR EACH ROW
EXECUTE FUNCTION view_route_tmp_delete();

CREATE FUNCTION view_route_perm_insert()
    RETURNS TRIGGER
AS $$
DECLARE id_ route.id%TYPE;
BEGIN
    INSERT INTO route (user_author_id, loc_dep, loc_arr, min_price)
    VALUES (new.user_author_id, new.loc_dep, new.loc_arr, new.min_price)
    RETURNING id INTO id_;
    INSERT INTO route_perm (id, even_week, odd_week, day_of_week, time_dep, time_arr)
    SELECT id_, new.even_week, new.odd_week, new.day_of_week, new.time_dep, new.time_arr;
    new.id := id_;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_perm_insert INSTEAD OF INSERT
    ON view_route_perm
    FOR EACH ROW
EXECUTE FUNCTION view_route_perm_insert();

CREATE FUNCTION view_route_perm_update()
    RETURNS TRIGGER
AS $$
BEGIN
    IF old.user_author_id != new.user_author_id THEN
        RAISE 'It is forbidden to update author of permanent route';
    END IF;
    UPDATE route SET loc_dep = new.loc_dep, loc_arr = new.loc_arr, min_price = new.min_price
    WHERE id = new.id AND user_author_id = new.user_author_id;
    UPDATE route_perm SET even_week = new.even_week, odd_week = new.odd_week, day_of_week = new.day_of_week,
                          time_dep = new.time_dep, time_arr = new.time_arr
    WHERE id = new.id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_perm_update INSTEAD OF UPDATE
    ON view_route_perm
    FOR EACH ROW
EXECUTE FUNCTION view_route_perm_update();

CREATE FUNCTION view_route_perm_delete()
    RETURNS TRIGGER
AS $$
BEGIN
    DELETE FROM route WHERE id = old.id;
    RETURN old;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER view_route_perm_delete INSTEAD OF DELETE
    ON view_route_perm
    FOR EACH ROW
EXECUTE FUNCTION view_route_perm_delete();

--CREATE INDEX ON user_ USING hash (vk_id); --TODO

--CREATE INDEX ON ad (user_author_id);
