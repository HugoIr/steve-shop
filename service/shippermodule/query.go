package shippermodule

const (
	addProductQuery = `
	INSERT INTO product (
		name,
		description,
		price,
		discount,
		stock
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) returning id
`
	getProductQuery = `
	SELECT
		name,
		description,
		max_weight,
		price,
		discount,
		stock
	FROM
		product
	WHERE
		id=$1
`

	getProductAllQuery = `
	SELECT
		*
	FROM
		product
`

	updateProductQuery = `
	UPDATE
		product
	SET
		name=$1,
		description=$2,
		price=$3,
		discount=$4,
		stock=$5
	WHERE
		id=$6
	returning id	
	
`

	removeProductQuery = `
	
	DELETE FROM
		product
	WHERE
		id=$1
`
)
