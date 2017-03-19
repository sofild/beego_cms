{{template "header.tpl" .}}
		
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">分类管理 / 分类添加</li>
			</ol>
		</div><!--/.row-->
						
		<div class="row">
			<div class="col-lg-12"><br /></div>			
		</div>

		<div class="row">
			<div class="col-lg-12">
				<div class="panel panel-default">
					<div class="panel-heading">分类编辑</div>
					<div class="panel-body">
						<div class="col-md-12">
							<form role="form" name="form1" id="form1" enctype="multipart/form-data" method="post" action="/articles/doadd">
							
								<div class="form-group">
									<label>分类名称</label>
									<input name="name" id="name" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>父分类</label>
									<select class="form-control" name="parent_id">
										<option value="0">请选择</option>
										<option value="2">分类 2</option>
										<option value="3">分类 3</option>
										<option value="4">分类 4</option>
									</select>
								</div>

								<div class="form-group">
									<label>排序</label>
									<input name="sort" id="sort" class="form-control" placeholder="">
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