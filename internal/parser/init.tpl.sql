-- fn: Insert
insert into {{.insert_table}} ({{.insert_columns}}) values ({{.insert_values}});
{{range .unique_indexes}}
-- fn: FindOneBy{{.UniqueNameJoin}}
select {{.SelectColumns}} from {{.Table}} where {{.WhereClause}} limit 1;
{{end}}
{{range .unique_indexes}}
-- fn: UpdateBy{{.UniqueNameJoin}}
update {{.Table}} set {{.UpdateSet}} where {{.WhereClause}};
{{end}}
{{range .unique_indexes}}
-- fn: DeleteBy{{.UniqueNameJoin}}
delete from {{.Table}} where {{.WhereClause}};
{{end}}