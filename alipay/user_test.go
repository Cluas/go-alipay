package alipay

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserService_InfoShare(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `
							{
								"alipay_user_info_share_response": {
									"code": "10000",
									"msg": "Success",
									"avatar": "https://tfs.alipayobjects.com/images/partner/T15ABtXk8bXXXXXXXX",
									"city": "杭州市",
									"gender": "m",
									"is_certified": "T",
									"is_student_certified": "F",
									"nick_name": "WinWen",
									"province": "浙江省",
									"user_id": "2088912161915762",
									"user_status": "T",
									"user_type": "2"
								}
							}`)
	})
	info, err := client.User.InfoShare(context.Background(), "token")
	if err != nil {
		t.Errorf("Account.GetBasicInfo returned error: %v", err)
	}
	want := &UserInfoShare{
		Avatar:             "https://tfs.alipayobjects.com/images/partner/T15ABtXk8bXXXXXXXX",
		City:               "杭州市",
		Gender:             "m",
		IsCertified:        "T",
		IsStudentCertified: "F",
		NickName:           "WinWen",
		Province:           "浙江省",
		UserID:             "2088912161915762",
		UserStatus:         "T",
		UserType:           "2",
	}
	if !reflect.DeepEqual(info, want) {
		t.Errorf("User.InfoShare returned %+v, want %+v", info, want)
	}

}
