{{template "header.tpl" .}}
			
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">文章管理 / 文章列表</li>
			</ol>
		</div><!--/.row-->
		
		<div class="row">
			<div class="col-lg-12">
				<br />
			</div>
		</div><!--/.row-->
				
		
		<div class="row">
			<div class="col-lg-12">
				<div class="panel panel-default">
					<div class="panel-heading">文章列表
						<span class="icon pull-right"><a href="/articles/add" title="新增文章"><em class="glyphicon glyphicon-s glyphicon-plus"></em></a></span>
					</div>
					<div class="panel-body">
						<table data-toggle="table" data-url="/articles/list"  data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-search="true" data-select-item-name="toolbar1" data-pagination="true" data-sort-name="name" data-sort-order="desc">
						    <thead>
						    <tr>
						        <th data-field="id" data-checkbox="true" >ID</th>
						        <th data-field="id" data-sortable="true">ID</th>
						        <th data-field="title">标题</th>
						        <th data-field="cate_name" data-sortable="true">分类</th>
						        <th data-field="imgUrl">图片</th>
						        <th data-field="time" data-sortable="true">更新时间</th>
						        <th data-field="op">操作</th>
						    </tr>
						    </thead>
						</table>
					</div>
				</div>
			</div>
		</div><!--/.row-->		
		
	</div><!--/.main-->
	<script src="/static/js/bootstrap-table.js"></script>
{{template "footer.tpl" .}}