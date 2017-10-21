package main

import (
	fmt "fmt"

	"github.com/keigodasu/grpc-ref-error-handle/search"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:8088"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := search.NewSearchServiceClient(conn)

	/********************************
	(人名="安倍" AND 年齢="22" AND 日付 btweeen 10/11 ~ 10/30) OR 状態=”未処理”
	********************************/
	filter01 := search.StructuredQuery_FieldFilter{
		Field:   &search.StructuredQuery_FieldReference{FieldPath: "人名"},
		Op:      search.StructuredQuery_FieldFilter_EQUAL,
		Keyword: "安倍",
	}

	filter02 := search.StructuredQuery_FieldFilter{
		Field:   &search.StructuredQuery_FieldReference{FieldPath: "年齢"},
		Op:      search.StructuredQuery_FieldFilter_EQUAL,
		Keyword: "22",
	}

	filter03 := search.StructuredQuery_FieldFilter{
		Field:   &search.StructuredQuery_FieldReference{FieldPath: "日付 "},
		Op:      search.StructuredQuery_FieldFilter_GREATER_THAN_OR_EQUAL,
		Keyword: "10/11",
	}

	filter04 := search.StructuredQuery_FieldFilter{
		Field:   &search.StructuredQuery_FieldReference{FieldPath: "日付 "},
		Op:      search.StructuredQuery_FieldFilter_LESS_THAN_OR_EQUAL,
		Keyword: "10/30",
	}

	filter001 := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_FieldFilter{
			FieldFilter: &filter01,
		},
	}

	filter002 := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_FieldFilter{
			FieldFilter: &filter02,
		},
	}

	filter003 := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_FieldFilter{
			FieldFilter: &filter03,
		},
	}

	filter004 := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_FieldFilter{
			FieldFilter: &filter04,
		},
	}

	comp := search.StructuredQuery_CompositeFilter{
		Op:      search.StructuredQuery_CompositeFilter_AND,
		Filters: []*search.StructuredQuery_Filter{&filter001, &filter002, &filter003, &filter004},
	}

	filter05 := search.StructuredQuery_FieldFilter{
		Field:   &search.StructuredQuery_FieldReference{FieldPath: "状態"},
		Op:      search.StructuredQuery_FieldFilter_LESS_THAN_OR_EQUAL,
		Keyword: "未処理",
	}

	filter005 := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_FieldFilter{
			FieldFilter: &filter05,
		},
	}

	filter := search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_CompositeFilter{
			CompositeFilter: &comp,
		},
	}

	comp02 := search.StructuredQuery_CompositeFilter{
		Op:      search.StructuredQuery_CompositeFilter_OR,
		Filters: []*search.StructuredQuery_Filter{&filter, &filter005},
	}

	filter = search.StructuredQuery_Filter{
		FilterType: &search.StructuredQuery_Filter_CompositeFilter{
			CompositeFilter: &comp02,
		},
	}

	strucQuery := search.StructuredQuery{
		Where: &filter,
	}

	ReqQuery := search.SearchRequest_StructuredQuery{
		StructuredQuery: &strucQuery,
	}

	req := search.SearchRequest{
		Query: &ReqQuery,
	}

	result, err := client.Search(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
