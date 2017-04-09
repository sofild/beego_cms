{{template "header.tpl" .}}
		
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">采集器管理 / 采集器编辑</li>
			</ol>
		</div><!--/.row-->
						
		<div class="row">
			<div class="col-lg-12"><br /></div>			
		</div>

		<div class="row">
			<div class="col-lg-12">
				<div class="panel panel-default">
					<div class="panel-heading">采集器编辑</div>
					<div class="panel-body">
						<div class="col-md-12">
							<form role="form" name="form1" id="form1" enctype="multipart/form-data" method="post" action="/spider/doadd">
							
								<div class="form-group">
									<label>名称</label>
									<input name="name" id="name" class="form-control" placeholder="">
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
									<label>链接</label>
									<input name="url" id="url" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>列表Tag</label>
									<input name="list" id="list" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>A标签Tag</label>
									<input name="link" id="link" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>标题Tag</label>
									<input name="title" id="title" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>图片Tag</label>
									<input name="pic" id="pic" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>作者Tag</label>
									<input name="author" id="author" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>内容Tag</label>
									<input name="content" id="content" class="form-control" placeholder="">
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
{{template "footer.tpl" .}}