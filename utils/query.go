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
)
