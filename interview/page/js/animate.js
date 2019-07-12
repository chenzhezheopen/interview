$(function(){
	$(".list1>li").click(function(){
		$(".list1>li").removeClass("active")
		$(this).addClass("active")
	})
	$(".mod li").click(function(){
//		alert("sss")
		$("#kes").text($(this).children("a").text())
	})
})
