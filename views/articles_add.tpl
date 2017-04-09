{{template "header.tpl" .}}
		
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">文章管理 / 文章添加</li>
			</ol>
		</div><!--/.row-->
						
		<div class="row">
			<div class="col-lg-12"><br /></div>			
		</div>

		<div class="row">
			<div class="col-lg-12">
				<div class="panel panel-default">
					<div class="panel-heading">文章编辑</div>
					<div class="panel-body">
						<div class="col-md-12">
							<form role="form" name="form1" id="form1" enctype="multipart/form-data" method="post" action="/articles/doadd">
							
								<div class="form-group">
									<label>标题</label>
									<input name="title" id="title" class="form-control" placeholder="" value="{{.Article.title}}">
								</div>
																
								<div class="form-group">
									<label>主图</label>
									<input type="file" name="upload" id="upload" runat="server" onChange="uploadImage()">
									<input type="hidden" name="pic" id="pic" value="{{.Article.pic}}" />
									<p class="help-block">jpg、gif、png  can be uploaded.</p>
									<p id="imgShow">
										{{if .Article.pic}}<img src="{{.Article.pic}}" width="100px" height="100px" />{{end}}
									</p>
								</div>
								
								<div class="form-group">
									<label>描述</label>
									<textarea class="form-control" rows="3" name="description" id="description">{{.Article.description}}</textarea>
								</div>

								<div class="form-group">
									<label>分类</label>
									{{$cateId := .Article.cate_id}}
									<select class="form-control" name="cate_id" data="{{$cateId}}">
									{{range $index,$value := .Cates}}
									{{$cid := $value.id}}
										<option value="{{$value.id}}">{{$value.name}}</option>
									{{end}}	
									</select>
								</div>

								<div class="form-group">
									<label>作者</label>
									<input name="author" id="author" class="form-control" placeholder="" value="{{.Content.author}}">
								</div>

								<div class="form-group">
									<label>来源</label>
									<input name="source" id="source" class="form-control" placeholder="" value="{{.Content.source}}">
								</div>

								<div class="form-group" style="height: 600px;">
									<label>内容</label>
									<textarea name="content" id="content" style="height:450px;">{{.Content.content}}</textarea>
								</div>

								<div class="form-group">
									<button type="submit" class="btn btn-primary">&nbsp;提&nbsp;&nbsp;交&nbsp;</button>
									<button type="reset" class="btn btn-default"> &nbsp;重&nbsp;&nbsp;置&nbsp;</button>
								</div>
							</div>
						</form>
					</div>
				</div>
			</div><!-- /.col-->
		</div><!-- /.row -->
		
	</div><!--/.main-->


	<script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.config.js"></script>
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/ueditor.all.min.js"> </script>
    <!--建议手动加在语言，避免在ie下有时因为加载语言失败导致编辑器加载失败-->
    <!--这里加载的语言文件会覆盖你在配置项目里添加的语言类型，比如你在配置项目里配置的是英文，这里加载的中文，那最后就是中文-->
    <script type="text/javascript" charset="utf-8" src="/static/ueditor/lang/zh-cn/zh-cn.js"></script>
    <script src="/static/js/jquery.form.js"></script>
    <script type="text/javascript">

    	var ue = UE.getEditor('content');


		function uploadImage() {
        	//判断是否有选择上传文件
            var imgPath = $("#upload").val();
            if (imgPath == "") {
                alert("请选择上传图片！");
                return;
            }
            console.log(imgPath)
            //判断上传文件的后缀名
            var strExtension = imgPath.substr(imgPath.lastIndexOf('.') + 1);
            if (strExtension != 'jpg' && strExtension != 'gif' && strExtension != 'png' && strExtension != 'bmp') {
                alert("请选择图片文件");
                return;
            }
            //构建上传的表单
            var fileInput = $("#upload").clone().prop("outerHTML");
            var formHtml = '<form id="uploadForm" name="uploadForm" method="post" enctype="multipart/form-data"><input type="file" name="upload" value="'+imgPath+'" width="0px" height="0px" /></form>';
            $("body").append(formHtml);
            console.log(formHtml);
            //
            //异步上传
            $("#uploadForm").ajaxSubmit({
                type: "POST",
                url: "/articles/upload",
                cache: false,
                success: function(data) {
                	$("#uploadForm").remove();
                	var status = parseInt(data.status)
                    if(status==1000){
                    	$("#pic").val(data.pic)
                    	$("#imgShow").html('<img src="'+data.pic+'" width="100px" height="100px" />')
                    	alert(data.msg)
                    }
                    else{
                    	alert(data.msg)
                    }
                },
                error: function(XMLHttpRequest, textStatus, errorThrown) {
                    alert("上传失败，请检查网络后重试");
                },
                async: true
            });
        }
    </script>
{{template "footer.tpl" .}}