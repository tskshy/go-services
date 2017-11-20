var G_HOST = "http://127.0.0.1:4243/"

$(document).ready(function(){
	init();
});

function init() {
/**
 * page init function, event binding, ...
 **/
	$("#nav-top div.nav-left").click(function() {
		$("#nav-left").toggle(500);
	});

	$("#nav-top i.search-top").click(function() {
		$("#search").slideToggle(1000);
	});

	$("#search-text").click(function() {
		$("#search-text").hide();
		$("#search-cancel").show();
	});

	$("#search-cancel").click(function() {
		$("#search-cancel").hide();
		$("#search-text").show();
		$("#nav-top i.search-top").click();
		$("#search-clear").click();
	});

	$("#search-clear").click(function() {
		$("#search-input").val("");
	});

	get_newest_list();
}

/*获取最新推荐*/
function get_newest_list() {
	for (var i = 0; i < 10; i++) {
		var node = $("#raw-model").clone(true).removeAttr("style").removeAttr("id");
		/*显示头像*/
		/*TODO*/

		/*显示昵称*/
		node.find("p").eq(0).text("昵称" + i);

		/*显示时间*/
		node.find("li").eq(0).text("2012-12-12 12:0" + i);

		/*显示标题*/
		node.find("a").eq(0).text("标题 " + i);

		/*显示简介*/
		node.find("p").eq(1).text("简介 " + i);

		node.find("div ul li").eq(1).text("点击量 " + i);
		node.find("div ul li").eq(2).text("点赞数 " + i);
		node.find("div ul li").eq(3).text("评论数 " + i);
		node.find("div ul li").eq(4).text("图文 " + i);

		$("#content-show").append(node);
	}
}
