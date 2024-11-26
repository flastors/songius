package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/flastors/songius/internal/music/storage"
	"github.com/flastors/songius/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	offset int
	fields []filter.Field
}

func NewFilterOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{
		limit:  options.Limit(),
		offset: options.Offset(),
		fields: options.Fields(),
	}
}

func (f *filterOptions) PaginationQuery() string {
	offsetQuery := ``
	if f.offset != 0 {
		offsetQuery = fmt.Sprintf(`OFFSET %d`, f.offset)
	}
	limitQuery := fmt.Sprintf(`LIMIT %d`, f.limit)
	q := fmt.Sprintf(`%s %s`, offsetQuery, limitQuery)
	return q
}

func (f *filterOptions) FilterQuery() (string, error) {
	q := ``
	if len(f.fields) != 0 {
		q = `WHERE`
		for i, field := range f.fields {
			if field.Type == filter.DataTypeStr {
				field.Value = "%" + field.Value + "%"
			}
			if field.Type == filter.DataTypeDate {
				inputLayout := "02.01.2006"
				outputLayout := "2006.01.02"
				if strings.Index(field.Value, ":") != -1 {
					dates := strings.Split(field.Value, ":")
					if len(dates) > 2 {
						return "", fmt.Errorf("too many dates, max 2")
					}
					for i, date := range dates {
						dateTime, err := time.Parse(inputLayout, date)
						if err != nil {
							return ``, fmt.Errorf("bad date format: %v", err)
						}
						dates[i] = dateTime.Format(outputLayout)
					}
					field.Value = fmt.Sprintf(`'%s' AND '%s'`, dates[0], dates[1])
					q += fmt.Sprintf(` %s %s %s`, field.Name, field.Operator, field.Value)
					if i+1 != len(f.fields) {
						q += ` AND`
					}
					continue
				} else {
					dateTime, err := time.Parse(inputLayout, field.Value)
					if err != nil {
						return ``, fmt.Errorf("bad date format: %v", err)
					}
					field.Value = dateTime.Format(outputLayout)
				}
			}
			q += fmt.Sprintf(` %s %s '%s'`, field.Name, field.Operator, field.Value)
			if i+1 != len(f.fields) {
				q += ` AND`
			}
		}
	}
	return q, nil
}
