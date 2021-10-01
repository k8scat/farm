package user

import "github.com/molizz/farm/model/db"

var (
	tableName = "farm_user"
	columns   = []string{
		"id",
		"thirdparty_id",
		"primary",
		"hash",
		"column_num_1",
		"column_num_2",
		"column_num_3",
		"column_num_4",
		"column_num_5",
		"column_num_6",
		"column_num_7",
		"column_num_8",
		"column_num_9",
		"column_num_10",
		"column_num_11",
		"column_num_12",
		"column_num_13",
		"column_num_14",
		"column_num_15",
		"column_num_16",
		"column_num_17",
		"column_num_18",
		"column_num_19",
		"column_num_20",
		"column_num_21",
		"column_num_22",
		"column_num_23",
		"column_num_24",
		"column_num_25",
		"column_num_26",
		"column_num_27",
		"column_num_28",
		"column_num_29",
		"column_num_30",
		"column_text_1",
		"column_text_2",
		"column_text_3",
		"column_text_4",
		"column_text_5",
		"column_text_6",
		"column_text_7",
		"column_text_8",
		"column_text_9",
		"column_text_10",
		"column_text_11",
		"column_text_12",
		"column_text_13",
		"column_text_14",
		"column_text_15",
		"column_text_16",
		"column_text_17",
		"column_text_18",
		"column_text_19",
		"column_text_20",
		"column_text_21",
		"column_text_22",
		"column_text_23",
		"column_text_24",
		"column_text_25",
		"column_text_26",
		"column_text_27",
		"column_text_28",
		"column_text_29",
		"column_text_30"}
)

type Queryer struct {
	db.Model
	Table   string
	Columns []string
}

// 三方用户表
type FarmUser struct {
	ID           uint    `db:"column:id"`
	ThirdpartyID uint    `db:"column:thirdparty_id"` //
	Primary      *string `db:"column:primary"`       //主键(通常由1或多个属性组成 )
	Hash         uint64  `db:"column:hash"`          // hash
	ColumnNum1   *int64  `db:"column:column_num_1"`
	ColumnNum2   *int64  `db:"column:column_num_2"`
	ColumnNum3   *int64  `db:"column:column_num_3"`
	ColumnNum4   *int64  `db:"column:column_num_4"`
	ColumnNum5   *int64  `db:"column:column_num_5"`
	ColumnNum6   *int64  `db:"column:column_num_6"`
	ColumnNum7   *int64  `db:"column:column_num_7"`
	ColumnNum8   *int64  `db:"column:column_num_8"`
	ColumnNum9   *int64  `db:"column:column_num_9"`
	ColumnNum10  *int64  `db:"column:column_num_10"`
	ColumnNum11  *int64  `db:"column:column_num_11"`
	ColumnNum12  *int64  `db:"column:column_num_12"`
	ColumnNum13  *int64  `db:"column:column_num_13"`
	ColumnNum14  *int64  `db:"column:column_num_14"`
	ColumnNum15  *int64  `db:"column:column_num_15"`
	ColumnNum16  *int64  `db:"column:column_num_16"`
	ColumnNum17  *int64  `db:"column:column_num_17"`
	ColumnNum18  *int64  `db:"column:column_num_18"`
	ColumnNum19  *int64  `db:"column:column_num_19"`
	ColumnNum20  *int64  `db:"column:column_num_20"`
	ColumnNum21  *int64  `db:"column:column_num_21"`
	ColumnNum22  *int64  `db:"column:column_num_22"`
	ColumnNum23  *int64  `db:"column:column_num_23"`
	ColumnNum24  *int64  `db:"column:column_num_24"`
	ColumnNum25  *int64  `db:"column:column_num_25"`
	ColumnNum26  *int64  `db:"column:column_num_26"`
	ColumnNum27  *int64  `db:"column:column_num_27"`
	ColumnNum28  *int64  `db:"column:column_num_28"`
	ColumnNum29  *int64  `db:"column:column_num_29"`
	ColumnNum30  *int64  `db:"column:column_num_30"`
	ColumnText1  *string `db:"column:column_text_1"`
	ColumnText2  *string `db:"column:column_text_2"`
	ColumnText3  *string `db:"column:column_text_3"`
	ColumnText4  *string `db:"column:column_text_4"`
	ColumnText5  *string `db:"column:column_text_5"`
	ColumnText6  *string `db:"column:column_text_6"`
	ColumnText7  *string `db:"column:column_text_7"`
	ColumnText8  *string `db:"column:column_text_8"`
	ColumnText9  *string `db:"column:column_text_9"`
	ColumnText10 *string `db:"column:column_text_10"`
	ColumnText11 *string `db:"column:column_text_11"`
	ColumnText12 *string `db:"column:column_text_12"`
	ColumnText13 *string `db:"column:column_text_13"`
	ColumnText14 *string `db:"column:column_text_14"`
	ColumnText15 *string `db:"column:column_text_15"`
	ColumnText16 *string `db:"column:column_text_16"`
	ColumnText17 *string `db:"column:column_text_17"`
	ColumnText18 *string `db:"column:column_text_18"`
	ColumnText19 *string `db:"column:column_text_19"`
	ColumnText20 *string `db:"column:column_text_20"`
	ColumnText21 *string `db:"column:column_text_21"`
	ColumnText22 *string `db:"column:column_text_22"`
	ColumnText23 *string `db:"column:column_text_23"`
	ColumnText24 *string `db:"column:column_text_24"`
	ColumnText25 *string `db:"column:column_text_25"`
	ColumnText26 *string `db:"column:column_text_26"`
	ColumnText27 *string `db:"column:column_text_27"`
	ColumnText28 *string `db:"column:column_text_28"`
	ColumnText29 *string `db:"column:column_text_29"`
	ColumnText30 *string `db:"column:column_text_30"`
}
