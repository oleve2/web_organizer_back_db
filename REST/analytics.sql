123

select * from activ_log
select * from activ_names
select * from activ_normative

-- base query
select *
from activ_log as T1
left join activ_names as T2 on T1.activ_name_id=T2.id
left join activ_normative as T3 on T1.activ_norm_id = T3.id



