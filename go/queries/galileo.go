package queries

var Items = `
SELECT * 
FROM data__atsgroup.__items 
where 
  has(tags,'VMWARECLUSTER@LAYER') or
  has(tags,'VMWAREGUEST@LAYER')
order by item_id
`
