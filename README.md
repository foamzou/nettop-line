# nettop-line
Better way to call nettop under delta mode as sub process on mac

## Usage
Execute the binary with some params. You can find the params with `nettop -h`

## For Example
When we wanna output delta `bytes_in and bytes_out` of every process per 1 second. Call that:

`./nettop-line -P -d -L 0 -J bytes_in,bytes_out -t external -s 1 -c`.

Then standard output continuously like that
```shell
,bytes_in,bytes_out,|SPLIT|kernel_task.0,3600,2347,|SPLIT|UserEventAgent.88,0,0,|SPLIT|logd.104,0,0,|SPLIT|apsd.122,0,0,|SPLIT|syspolicyd.154,0,0,|SPLIT|mDNSResponder.224,2612,1227,|SPLIT|SubmitDiagInfo.292,0,0,|SPLIT|rapportd.345,0,0,|SPLIT|goland.410,0,0,|SPLIT|WeChat.420,0,0,|SPLIT|Spark.423,0,0,|SPLIT|SystemUIServer.432,0,0,|SPLIT|Google Chrome H.487,424,464,|SPLIT|Slack Helper.516,0,0,|SPLIT|corespeechd.2667,0,0,|SPLIT|SogouServices.2669,0,0,|SPLIT|NeteaseMusic.4137,0,0,|SPLIT|com.apple.WebKi.4139,0,0,|SPLIT|PowerChime.4776,0,0,|SPLIT|DingTalk.56492,0,0,|SPLIT|Postman Helper.70125,0,0,|SPLIT|netbiosd.27658,0,0,|SPLIT|ss-local.37764,0,0,|SPLIT|findmydeviced.39862,0,0,|SPLIT|webstorm.41606,0,0,|SPLIT|biometrickitd.2449,0,0,|SPLIT|racoon.3086,0,0,
,bytes_in,bytes_out,|SPLIT|kernel_task.0,63,240,|SPLIT|UserEventAgent.88,0,0,|SPLIT|logd.104,0,0,|SPLIT|apsd.122,0,0,|SPLIT|syspolicyd.154,0,0,|SPLIT|mDNSResponder.224,0,87,|SPLIT|SubmitDiagInfo.292,0,0,|SPLIT|rapportd.345,0,0,|SPLIT|goland.410,0,0,|SPLIT|WeChat.420,0,0,|SPLIT|Spark.423,0,0,|SPLIT|SystemUIServer.432,0,0,|SPLIT|Google Chrome H.487,25,33,|SPLIT|Slack Helper.516,0,56,|SPLIT|corespeechd.2667,0,0,|SPLIT|SogouServices.2669,0,0,|SPLIT|NeteaseMusic.4137,0,0,|SPLIT|com.apple.WebKi.4139,0,0,|SPLIT|PowerChime.4776,0,0,|SPLIT|DingTalk.56492,0,0,|SPLIT|Postman Helper.70125,0,0,|SPLIT|netbiosd.27658,0,0,|SPLIT|ss-local.37764,0,0,|SPLIT|findmydeviced.39862,0,0,|SPLIT|webstorm.41606,0,0,|SPLIT|biometrickitd.2449,0,0,|SPLIT|racoon.3086,0,0,
```

## Build
```
go build -ldflags="-w -s"
# We can use upx to compress our bin file, although it would cost some memory while our target program running
upx -3 -o nettop-line-3 nettop-line && mv nettop-line-3 nettop-line 
```

## Notice
The program would drop the first output from nettop which was "dirty data"(For it was not really increase data).

So, You do not need to drop the dirty data again

## Why wrap nettop
If we want to get delta in/out traffic per process in our program, we can call `nettop`.

But there are some troubled while calling it as sub process directly:
- nettop would cause high cpu usage if not block stdin of nettop. It's seem bug of nettop
- stdout process line by line, we can't know which line is end of the duration

The program by wrap nettop, avoid cpu issue AND output result per duration line by line. You can split process result by split `|SPLIT|`
