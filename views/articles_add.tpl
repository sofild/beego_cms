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
									<input name="title" id="title" class="form-control" placeholder="">
								</div>
																
								<div class="form-group">
									<label>主图</label>
									<input type="file" name="pic">
									 <p class="help-block">jpg、gif、png  can be uploaded.</p>
								</div>
								
								<div class="form-group">
									<label>描述</label>
									<textarea class="form-control" rows="3" name="description" id="description"></textarea>
								</div>

								<div class="form-group">
									<label>分类</label>
									<select class="form-control" name="cate_id">
										<option value="1">分类 1</option>
										<option value="2">分类 2</option>
										<option value="3">分类 3</option>
										<option value="4">分类 4</option>
									</select>
								</div>

								<div class="form-group">
									<label>作者</label>
									<input name="author" id="author" class="form-control" placeholder="">
								</div>

								<div class="form-group">
									<label>来源</label>
									<input name="source" id="source" class="form-control" placeholder="">
								</div>

								<div class="form-group" style="height: 600px;">
									<label>内容</label>
									<textarea name="content" id="content" style="height:450px;"></textarea>
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
    <script type="text/javascript">

    	var ue = UE.getEditor('content');



    </script>
{{template "footer.tpl" .}}