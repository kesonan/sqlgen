insert into {{.insert_table}} ({{.insert_columns}}) values ({{.insert_values}});
{{range .unique_indexes}}
select {{.SelectColumns}} from {{.Table}} where {{.WhereClause}} limit 1,20;
{{end}}
{{range .unique_indexes}}
update {{.Table}} set {{.UpdateSet}} where {{.WhereClause}};
{{end}}
{{range .unique_indexes}}
delete from {{.Table}} where {{.WhereClause}};
{{end}}