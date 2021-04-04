package date_test

import "testing"
import "date"

option now = () => (2030-01-01T00:00:00Z)

inData = "
#datatype,string,long,dateTime:RFC3339,string,string,double
#group,false,false,false,true,true,false
#default,_result,,,,,
,result,table,_time,_measurement,_field,_value
,,0,2018-05-01T19:53:00Z,_m,FF,1
,,0,2018-05-02T19:53:10Z,_m,FF,1
,,0,2018-05-03T19:53:20Z,_m,FF,1
,,0,2018-05-04T19:53:30Z,_m,FF,1
,,0,2018-05-05T19:53:40Z,_m,FF,1
,,0,2018-05-06T19:53:50Z,_m,FF,1
,,1,2018-05-07T19:53:00Z,_m,QQ,1
,,1,2018-05-08T19:53:10Z,_m,QQ,1
,,1,2018-05-09T19:53:20Z,_m,QQ,1
,,1,2018-05-10T19:53:30Z,_m,QQ,1
,,1,2018-05-11T19:53:40Z,_m,QQ,1
,,1,2018-05-12T19:53:50Z,_m,QQ,1
,,1,2018-05-13T19:54:00Z,_m,QQ,1
,,1,2018-05-14T19:54:10Z,_m,QQ,1
,,1,2018-05-15T19:54:20Z,_m,QQ,1
,,2,2018-05-16T19:53:00Z,_m,RR,1
,,2,2018-05-17T19:53:10Z,_m,RR,1
,,2,2018-05-18T19:53:20Z,_m,RR,1
,,2,2018-05-19T19:53:30Z,_m,RR,1
,,3,2018-05-20T19:53:40Z,_m,SR,1
,,3,2018-05-21T19:53:50Z,_m,SR,1
,,3,2018-05-22T19:54:00Z,_m,SR,1
,,3,2018-05-23T19:54:00Z,_m,SR,1
,,3,2018-05-24T19:54:00Z,_m,SR,1
,,3,2018-05-25T19:54:00Z,_m,SR,1
,,3,2018-05-26T19:54:00Z,_m,SR,1
,,3,2018-05-27T19:54:00Z,_m,SR,1
,,3,2018-05-28T19:54:00Z,_m,SR,1
,,3,2018-05-29T19:54:00Z,_m,SR,1
,,3,2018-05-30T19:54:00Z,_m,SR,1
,,3,2018-05-31T19:54:00Z,_m,SR,1
"

outData = "
#group,false,false,true,true,true,true,false,false
#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,dateTime:RFC3339,long
#default,_result,,,,,,,
,result,table,_start,_stop,_field,_measurement,_time,_value
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-01T19:53:00Z,1
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-02T19:53:10Z,2
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-03T19:53:20Z,3
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-04T19:53:30Z,4
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-05T19:53:40Z,5
,,0,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,FF,_m,2018-05-06T19:53:50Z,6
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-07T19:53:00Z,7
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-08T19:53:10Z,8
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-09T19:53:20Z,9
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-10T19:53:30Z,10
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-11T19:53:40Z,11
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-12T19:53:50Z,12
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-13T19:54:00Z,13
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-14T19:54:10Z,14
,,1,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,QQ,_m,2018-05-15T19:54:20Z,15
,,2,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,RR,_m,2018-05-16T19:53:00Z,16
,,2,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,RR,_m,2018-05-17T19:53:10Z,17
,,2,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,RR,_m,2018-05-18T19:53:20Z,18
,,2,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,RR,_m,2018-05-19T19:53:30Z,19
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-20T19:53:40Z,20
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-21T19:53:50Z,21
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-22T19:54:00Z,22
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-23T19:54:00Z,23
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-24T19:54:00Z,24
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-25T19:54:00Z,25
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-26T19:54:00Z,26
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-27T19:54:00Z,27
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-28T19:54:00Z,28
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-29T19:54:00Z,29
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-30T19:54:00Z,30
,,3,2018-01-01T00:00:00Z,2030-01-01T00:00:00Z,SR,_m,2018-05-31T19:54:00Z,31
"

t_time_month_day = (table=<-) =>
	(table
	    |> range(start: 2018-01-01T00:00:00Z)
		|> map(fn: (r) => ({r with _value: date.monthDay(t: r._time)})))

test _time_month_day = () =>
	({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_time_month_day})
