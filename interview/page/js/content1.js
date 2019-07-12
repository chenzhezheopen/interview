//new Vue({
//  el: "#app-4",
//  data: {
//      images: [
//          { src: "./image/1.jpg", href: "#1" },
//          { src: "./image/2.jpg", href: "#2" },
//          { src: "./image/3.jpg", href: "#3" },
//          { src: "./image/4.jpg", href: "#4" },
//          { src: "./image/5.jpg", href: "#5" }
//      ],
//      config: {
//          effect: "slide",
//          dot: "#custom-dot"
//      }
//  }
//});
$(function(){
	$("#history").click(function(){
		
	})
	$("#change button").click(function(){
//		alert("aaa")
		$("#change button").removeClass("active2")
		$(this).addClass("active2")
	})
	$("#change button:nth-child(1)").click(function(){
		$("#form").css("display","block")
		$("#form1").css("display","none")
	})
	$("#change button:nth-child(2)").click(function(){
//		alert("1")
		$("#form1").css("display","block")
		$("#form").css("display","none")
	})
	$("#history").click(function(){
//		alert("1")
		$('[id^="content"]').hide(100)
		$("#content2").show(600)
	})
	$("#an_content1").click(function(){
		$('[id^="content"]').hide(300)
		$("#content1").show(600)
	})
	
	$("#an_content3").click(function(){
		$("#content3>.row").empty()
		$('[id^="content"]').hide(300)
		$("#content3").show(600)
		console.log("收到消息后触发 message received:aaaaaaaa ");
		var wsurl = "ws://localhost:5656/user_obt"
		sock = new WebSocket(wsurl);
		sock.onmessage = function(e) {
				var data=JSON.parse(e.data);
	            console.log("收到消息后触发 message received: " + e.data);
	//          console.log("收到消息后触发 message received: " + data.Value.Name);
				var newli=$("<div class='col-sm-6 col-md-4'><div class='thumbnail'><img src='image/"+data.Value.Img+"' ><div class='caption'><h3>"+data.Value.Name+"</h3><p>"+data.Value.Occ+"</p><p><a href='#' class='btn btn-primary amin' role='button'>咨询</a> <a href='#' class='btn btn-default' role='button'>信息</a></p></div></div></div>")
				$("#content3>.row").append(newli)
//				var data=""
				
	    }
		
	})
//	$(".amin").on('click', function(){
//	 alert('OK');
//	});
//	$("#communication").siblings().click(function(){
//		$("#communication").hide(100)
//	})
})
