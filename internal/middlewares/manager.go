package middlewares

type Manager struct {
	RecoverMiddleware *RecoverMiddleware
	AuthMiddleware    *AuthMiddleware
}

func NewManager(recoverMiddleware *RecoverMiddleware, authMiddleware *AuthMiddleware) *Manager {
	return &Manager{
		RecoverMiddleware: recoverMiddleware,
		AuthMiddleware:    authMiddleware,
	}
}
