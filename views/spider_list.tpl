{{template "header.tpl" .}}
			
	<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">			
		<div class="row">
			<ol class="breadcrumb">
				<li><a href="#"><span class="glyphicon glyphicon-home"></span></a></li>
				<li class="active">采集器管理 / 采集器列表</li>
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
					<div class="panel-heading">
						采集器列表
						<span class="icon pull-right"><a href="/spider/add" title="新增配置"><em class="glyphicon glyphicon-s glyphicon-plus"></em></a></span>	
					</div>
					<div class="panel-body">
						<table data-toggle="table" data-url="/spider/list"  data-show-refresh="true" data-show-toggle="true" data-show-columns="true" data-search="true" data-select-item-name="toolbar1" data-pagination="true" data-sort-name="name" data-sort-order="desc">
						    <thead>
						    <tr>
						        <th data-field="id" data-checkbox="true" >ID</th>
						        <th data-field="name" data-sortable="true">名称</th>
						        <th data-field="cate_name">分类名称</th>
						        <th data-field="url" data-sortable="true">链接</th>
						        <th data-field="time" data-sortable="true">添加时间</th>
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