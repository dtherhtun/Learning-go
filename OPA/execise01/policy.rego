package rules

default allow = false

users := {
	"alice": {"manager": "charlie", "title": "salesperson"},
	"bob": {"manager": "charlie", "title": "salesperson"},
	"charlie": {"manager": "dave", "title": "manager"},
	"dave": {"manager": null, "title": "ceo"},
}

allow {
	input.path == ["cars"]
	input.method = "GET"
}


test_car_read_positive {
    in = {
       "method": "GET",
       "path": ["cars"],
       "user": "alice"
    }
    allow == true with input as in
}

test_car_read_negative {
    in = {
       "method": "GET",
       "path": ["nonexistent"],
       "user": "alice"
    }
    allow == false with input as in
}

user_is_employee {
    users[input.user]
}

user_is_manager  {
    users[input.user].title != "salesperson"
}

allow {
    input.path = ["cars"]
    input.method = "POST"
    user_is_manager
}

allow {
    user_is_employee
    input.method == "GET"
    input.path == ["cars", carid]
}

test_car_create_negative {
    in = {
       "method": "POST",
       "path": ["cars"],
       "user": "alice"
    }
    allow == false with input as in
}

test_car_create_positive {
    in = {
       "method": "POST",
       "path": ["cars"],
       "user": "charlie"
    }
    allow == true with input as in
}