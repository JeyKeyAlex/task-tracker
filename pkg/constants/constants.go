package constants

const (
	Schema string = `
			CREATE TABLE IF NOT EXISTS scheduler (
    			id INTEGER PRIMARY KEY AUTOINCREMENT,
    			date CHAR(8) NOT NULL DEFAULT "",
    			title VARCHAR(255) NOT NULL DEFAULT "",
    			comment TEXT,
    			repeat VARCHAR(128) NOT NULL DEFAULT ""
			);

			CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`
)

const (
	DataFormat    string = "20060102"
	DaySign       string = "d"
	MonthSign     string = "m"
	WeekSign      string = "w"
	YearSign      string = "y"
	DaysMinValue  int    = 1
	DaysMaxValue  int    = 400
	WeeksMinValue int    = 1
	WeeksMaxValue int    = 7
)
