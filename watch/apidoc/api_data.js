define({ "api": [
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "Init",
    "group": "watch",
    "description": "<p>调用runtime.GOOS判断系统类型，从而确定注释符</p>",
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  },
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "listen",
    "group": "watch",
    "description": "<p>监听文件是否被修改的函数</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  },
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "getFileMes",
    "group": "watch",
    "description": "<p>获取文件内容</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "map[string]string",
            "optional": false,
            "field": "file_mes",
            "description": "<p>文件内容</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  },
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "read_file",
    "group": "watch",
    "description": "<p>读取文件并返回map[string]string和自定义错误类型</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "map[string]string",
            "optional": false,
            "field": "file_mes",
            "description": "<p>文件内容</p>"
          },
          {
            "group": "Success 200",
            "type": "error_mes",
            "optional": false,
            "field": "err",
            "description": "<p>自定义错误信息</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  },
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "Watch",
    "group": "watch",
    "description": "<p>读取、打印、监听文件，如果文件内容被修改则输出修改后的文件内容</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "file_name",
            "description": "<p>文件名</p>"
          },
          {
            "group": "Parameter",
            "type": "Listener",
            "optional": false,
            "field": "Listen",
            "description": "<p>监听器</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "map[string]string",
            "optional": false,
            "field": "file_mes",
            "description": "<p>文件内容</p>"
          },
          {
            "group": "Success 200",
            "type": "error_mes",
            "optional": false,
            "field": "err",
            "description": "<p>自定义错误信息</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  },
  {
    "type": "方法",
    "url": "./watch/watch.go",
    "title": "print_config",
    "group": "watch",
    "description": "<p>输出相应config</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "map[string]string",
            "optional": false,
            "field": "config",
            "description": "<p>保存配置文件的对应关系</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "watch/watch.go",
    "groupTitle": "watch",
    "name": "方法WatchWatchGo"
  }
] });
