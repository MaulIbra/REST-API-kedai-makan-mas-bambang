package utils

const(
	SELECT_MENU = `SELECT 
	p.menu_id,
    p.menu_name,
    p.stok,
	p.price,
    p.menu_active,
	COALESCE(c.category_id, '') as category_id ,
    COALESCE(c.category_name, '') as category_name
	FROM m_menu p left join m_category c on p.category_id = c.category_id
	WHERE p.menu_active = 1`
	SELECT_MENU_BY_ID = `SELECT 
	p.menu_id,
    p.menu_name,
    p.stok,
	p.price,
    p.menu_active,
	COALESCE(c.category_id, '') as category_id ,
    COALESCE(c.category_name, '') as category_name
	FROM m_menu p left join m_category c on p.category_id = c.category_id
	WHERE p.menu_id = ?`
	INSERT_MENU = `INSERT INTO m_menu values(?,?,?,?,?,?)`
	UPDATE_MENU = `UPDATE m_menu SET category_id=?,menu_name=?,stok=?,price=?,menu_active=? where menu_id=?;`
	DELETE_MENU = `UPDATE m_menu set menu_active=? where menu_id=?`

	INSERT_TRANSACTION = `INSERT INTO m_transaction values(?,?,?,?)`
	SELECT_TRANSACTION_BY_ID = `SELECT 
	t.transaction_date,
	t.transaction_id,
    t.menu_id,
    m.menu_name,
    t.quantity,
    m.price,
    sum(t.quantity * m.price) as total_price
	FROM m_transaction t inner join m_menu m on t.menu_id = m.menu_id 
    where t.transaction_id = ?
	group by t.transaction_id,t.menu_id`
	SELECT_TRANSACTION = `SELECT 
	t.transaction_date,
	t.transaction_id,
    t.menu_id,
    m.menu_name,
    t.quantity,
    m.price,
    sum(t.quantity * m.price) as total_price
	FROM m_transaction t inner join m_menu m on t.menu_id = m.menu_id 
    where t.transaction_date like ?
	group by t.transaction_id,t.menu_id
	order by t.transaction_id
	`
	SELECT_STOCK_MENU = `Select stok,menu_name from m_menu where menu_id=? and menu_active=1`
	UPDATE_STOCK_MENU = `UPDATE m_menu set stok=? where menu_id = ?`
	INSERT_ADDITIONAL_SERVICE_IN_TRANSACTION = `INSERT INTO m_transaction_has_m_additional values(?,?)`
	SELECT_ADDITIONAL_SERVICE_IN_TRANSACTION = `SELECT a.additional_id,a.additional_name,a.additional_price FROM m_transaction_has_m_additional tm join m_additional a on tm.additional_id=a.additional_id and tm.transaction_id=?`
)
