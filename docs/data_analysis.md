# Data analysis

The CSV files can be used to perform data analysis in many ways, For example

* Load to a spreadsheet like Microsoft Excel
* Load to Elasticsearch using Filebeat and visualize with Kibana
* Use python
etc.

This blog post does data analysis with python. http://www.markwk.com/data-analysis-for-apple-health.html in jupyter notebook. 

* Delete the first cell as we are not using python to convert health data to csv.
* Replace the timezone with your one in the following line:

```
convert_tz = lambda x: x.to_pydatetime().replace(tzinfo=pytz.utc).astimezone(pytz.timezone('Asia/Shanghai'))
```
For example for US East coast:
```
convert_tz = lambda x: x.to_pydatetime().replace(tzinfo=pytz.utc).astimezone(pytz.timezone('America/new_york'))
```

All the steps worked successfully for me.
