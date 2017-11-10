var G_HOST = "http://127.0.0.1:4243/"

$(document).ready(function(){
	init();
});

function init() {
/**
 * page init function
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
}
