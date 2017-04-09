
	<div id="sidebar-collapse" class="col-sm-3 col-lg-2 sidebar">
		<form role="search">
			<div class="form-group">
				<input type="text" class="form-control" placeholder="Search">
			</div>
		</form>
		<ul class="nav menu">
			<li {{ if eq .MenuName "article" }}class="active"{{ end }}><a href="/articles"><span class="glyphicon glyphicon-dashboard"></span> 文章管理</a></li>
			<li {{ if eq .MenuName "cate" }}class="active"{{ end }}><a href="/cate"><span class="glyphicon glyphicon-th"></span> 分类管理</a></li>
			<li {{ if eq .MenuName "spider" }}class="active"{{ end }}><a href="/spider"><span class="glyphicon glyphicon-stats"></span> 采集器管理</a></li>
			<!--
			<li><a href="tables.html"><span class="glyphicon glyphicon-list-alt"></span> Tables</a></li>
			<li><a href="forms.html"><span class="glyphicon glyphicon-pencil"></span> Forms</a></li>
			<li><a href="panels.html"><span class="glyphicon glyphicon-info-sign"></span> Alerts &amp; Panels</a></li>
			<li class="parent ">
				<a href="#">
					<span class="glyphicon glyphicon-list"></span> Dropdown <span data-toggle="collapse" href="#sub-item-1" class="icon pull-right"><em class="glyphicon glyphicon-s glyphicon-plus"></em></span> 
				</a>
				<ul class="children collapse" id="sub-item-1">
					<li>
						<a class="" href="#">
							<span class="glyphicon glyphicon-share-alt"></span> Sub Item 1
						</a>
					</li>
					<li>
						<a class="" href="#">
							<span class="glyphicon glyphicon-share-alt"></span> Sub Item 2
						</a>
					</li>
					<li>
						<a class="" href="#">
							<span class="glyphicon glyphicon-share-alt"></span> Sub Item 3
						</a>
					</li>
				</ul>
			</li>
			-->
			<li role="presentation" class="divider"></li>
		</ul>
		<div class="attribution"></div>
	</div><!--/.sidebar-->