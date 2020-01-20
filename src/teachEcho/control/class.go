package control

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"strconv"
	"teachEcho/model"
)

func ClassAll(ctx echo.Context) error {
	mods, err := model.ClassAll()
	if err != nil {
		ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", mods))
}

func ClassPage(ctx echo.Context) error {
	pi, err := strconv.Atoi(ctx.FormValue("page"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页索引错误", err.Error()))
	}
	if pi < 1 {
		return ctx.JSON(utils.ErrIpt("分页索引范围错误", err.Error()))
	}
	ps, err := strconv.Atoi(ctx.FormValue("limit"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页大小错误", err.Error()))
	}
	if ps < 1 || ps > 50 {
		return ctx.JSON(utils.ErrIpt("分页大小范围错误", err.Error()))
	}
	count := model.ClassCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ClassPage(pi, ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.PageLayUi("分类分页数据", mods, count))
}
