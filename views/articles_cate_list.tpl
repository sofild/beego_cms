{{template "header.tpl" .}}
			
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">分类管理 / 分类列表</li>
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
					<div class="panel-heading">分类列表</div>
					<div class="panel-body">
						<table data-toggle="table" data-url="/cate/data"  data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-search="true" data-select-item-name="toolbar1" data-pagination="true" data-sort-name="name" data-sort-order="desc">
						    <thead>
						    <tr>
						        <th data-field="id" data-checkbox="true" >ID</th>
						        <th data-field="id" data-sortable="true">ID</th>
						        <th data-field="parent_id">父分类</th>
						        <th data-field="name" data-sortable="true">分类名称</th>
						        <th data-field="sort">排序</th>
						        <th data-field="time" data-sortable="true">添加时间</th>
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