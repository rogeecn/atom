package providers

import (
	_ "atom/providers/captcha"
	_ "atom/providers/captcha/driver"
	_ "atom/providers/config"
	_ "atom/providers/database"
	_ "atom/providers/faker"
	_ "atom/providers/http"
	_ "atom/providers/jwt"
	_ "atom/providers/log"
	_ "atom/providers/query"
	_ "atom/providers/single_flight"
	_ "atom/providers/uuid"
)
