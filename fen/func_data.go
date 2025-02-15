package fen

import (
	"github.com/gofiber/fiber/v3"
)

func DataFunc0[T any](
	f func(fiber.Ctx) (T, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		data, err := f(ctx)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc1[T, P1 any](
	f func(fiber.Ctx, P1) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p, err := pf1(ctx)
		if err != nil {
			return err
		}

		data, err := f(ctx, p)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc2[T, P1, P2 any](
	f func(fiber.Ctx, P1, P2) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc3[T, P1, P2, P3 any](
	f func(fiber.Ctx, P1, P2, P3) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc4[T, P1, P2, P3, P4 any](
	f func(fiber.Ctx, P1, P2, P3, P4) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc5[T, P1, P2, P3, P4, P5 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc6[T, P1, P2, P3, P4, P5, P6 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5, P6) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
	pf6 func(fiber.Ctx) (P6, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		p6, err := pf6(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc7[T, P1, P2, P3, P4, P5, P6, P7 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5, P6, P7) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
	pf6 func(fiber.Ctx) (P6, error),
	pf7 func(fiber.Ctx) (P7, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		p6, err := pf6(ctx)
		if err != nil {
			return err
		}
		p7, err := pf7(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc8[T, P1, P2, P3, P4, P5, P6, P7, P8 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5, P6, P7, P8) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
	pf6 func(fiber.Ctx) (P6, error),
	pf7 func(fiber.Ctx) (P7, error),
	pf8 func(fiber.Ctx) (P8, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		p6, err := pf6(ctx)
		if err != nil {
			return err
		}
		p7, err := pf7(ctx)
		if err != nil {
			return err
		}
		p8, err := pf8(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc9[T, P1, P2, P3, P4, P5, P6, P7, P8, P9 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5, P6, P7, P8, P9) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
	pf6 func(fiber.Ctx) (P6, error),
	pf7 func(fiber.Ctx) (P7, error),
	pf8 func(fiber.Ctx) (P8, error),
	pf9 func(fiber.Ctx) (P9, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		p6, err := pf6(ctx)
		if err != nil {
			return err
		}
		p7, err := pf7(ctx)
		if err != nil {
			return err
		}
		p8, err := pf8(ctx)
		if err != nil {
			return err
		}
		p9, err := pf9(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func DataFunc10[T, P1, P2, P3, P4, P5, P6, P7, P8, P9, P10 any](
	f func(fiber.Ctx, P1, P2, P3, P4, P5, P6, P7, P8, P9, P10) (T, error),
	pf1 func(fiber.Ctx) (P1, error),
	pf2 func(fiber.Ctx) (P2, error),
	pf3 func(fiber.Ctx) (P3, error),
	pf4 func(fiber.Ctx) (P4, error),
	pf5 func(fiber.Ctx) (P5, error),
	pf6 func(fiber.Ctx) (P6, error),
	pf7 func(fiber.Ctx) (P7, error),
	pf8 func(fiber.Ctx) (P8, error),
	pf9 func(fiber.Ctx) (P9, error),
	pf10 func(fiber.Ctx) (P10, error),
) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		p1, err := pf1(ctx)
		if err != nil {
			return err
		}
		p2, err := pf2(ctx)
		if err != nil {
			return err
		}
		p3, err := pf3(ctx)
		if err != nil {
			return err
		}
		p4, err := pf4(ctx)
		if err != nil {
			return err
		}
		p5, err := pf5(ctx)
		if err != nil {
			return err
		}
		p6, err := pf6(ctx)
		if err != nil {
			return err
		}
		p7, err := pf7(ctx)
		if err != nil {
			return err
		}
		p8, err := pf8(ctx)
		if err != nil {
			return err
		}
		p9, err := pf9(ctx)
		if err != nil {
			return err
		}
		p10, err := pf10(ctx)
		if err != nil {
			return err
		}
		data, err := f(ctx, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
