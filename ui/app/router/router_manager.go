package router

func (r *Router) getClickedNaviBtnIndex(gtx LC) int {
	// 获得被点击的导航按钮的索引
	for i, c := range r.NaviMenuBarComp.MenuButtonsComp {
		if c.Clickable.Clicked(gtx) {
			return i
		}
	}
	return -1 // 没有按钮被点击则返回-1
}

func (r *Router) updateSelectedPageIndex(gtx LC) {
	// 当有导航按钮被点击后,更新当前正浏览的页面索引
	clickedIndex := r.getClickedNaviBtnIndex(gtx)
	if clickedIndex != -1 {
		r.CurrSelectedPageIndex = clickedIndex
	}
}
