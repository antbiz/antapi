<div lang="zh-cn" dir="ltr" class="mw-content-ltr"><h2 class="wiki_title"> <span class="mw-headline" id=".E9.94.99.E8.AF.AF.E4.BB.A3.E7.A0.81.E8.AF.B4.E6.98.8E">错误代码说明</span></h2>
<h3> <span class="mw-headline" id=".E9.94.99.E8.AF.AF.E8.BF.94.E5.9B.9E.E5.80.BC.E6.A0.BC.E5.BC.8F">错误返回值格式</span></h3>
<p>JSON
</p>
<pre>{
	"request"&nbsp;: "/statuses/home_timeline.json",
	"error_code"&nbsp;: "20502",
	"error"&nbsp;: "Need you follow uid."
}
</pre>
<h3> <span class="mw-headline" id=".E9.94.99.E8.AF.AF.E4.BB.A3.E7.A0.81.E8.AF.B4.E6.98.8E_2">错误代码说明</span></h3>
<p>20502
</p>
<table border="1" cellspacing="0" cellpadding="0" width="100%" class="parameters" style="border-color: #CCCCCC;">
  <tbody><tr>
    <td width="40%" style="text-align:left;padding-left:5px;border:1px solid #cccccc">2</td>
    <td width="30%" style="text-align:left;padding-left:5px;border:1px solid #cccccc">05</td>
    <td width="30%" style="text-align:left;padding-left:5px;border:1px solid #cccccc">02</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">服务级错误（1为系统级错误）</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">服务模块代码</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">具体错误代码</td>
  </tr>
</tbody></table>
<p><br>
</p>
<h2 class="wiki_title"> <span class="mw-headline" id=".E9.94.99.E8.AF.AF.E4.BB.A3.E7.A0.81.E5.AF.B9.E7.85.A7.E8.A1.A8">错误代码对照表</span></h2>
<h3> <span class="mw-headline" id=".E7.B3.BB.E7.BB.9F.E7.BA.A7.E9.94.99.E8.AF.AF.E4.BB.A3.E7.A0.81">系统级错误代码</span></h3>
<table border="1" cellspacing="0" cellpadding="0" width="100%" class="parameters" style="border-color: #CCCCCC;">
  <tbody><tr>
    <td width="10%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">错误代码</td>
    <td width="55%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">错误信息</td>
    <td width="35%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">详细描述</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10001</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">System error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">系统错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10002</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Service unavailable</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">服务暂停</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10003</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Remote service error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">远程服务错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10004</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP限制不能请求该资源</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10005</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Permission denied, need a high level appkey</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">该资源需要appkey拥有授权</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10006</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Source paramter (appkey) is missing</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">缺少source (appkey) 参数</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10007</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Unsupport mediatype (%s)</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不支持的MediaType (%s)</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10008</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Param error, see doc for more info</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数错误，请参考API文档</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10009</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Too many pending tasks, system is busy</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">任务过多，系统繁忙</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10010</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Job expired</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">任务超时</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10011</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">RPC error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">RPC错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10012</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal request</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">非法请求</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10013</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Invalid weibo user</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不合法的微博用户</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10014</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Insufficient app permissions</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">应用的接口访问权限受限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10016</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Miss required parameter (%s) , see doc for more info</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">缺失必选参数 (%s)，请参考API文档</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10017</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Parameter (%s)'s value invalid, expect (%s) , but get (%s) , see doc for more info</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数值非法，需为 (%s)，实际为 (%s)，请参考API文档</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10018</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Request body length over limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">请求长度超过限制</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10020</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Request api not found</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">接口不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10021</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">HTTP method is not suported for this request</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">请求的HTTP METHOD不支持，请检查是否选择了正确的POST/GET方式
</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10022</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP requests out of rate limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP请求频次超过上限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10023</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">User requests out of rate limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">用户请求频次超过上限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">10024</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">User requests for (%s) out of rate limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">用户请求特殊接口 (%s) 频次超过上限</td>
  </tr>
</tbody></table>
<h3> <span class="mw-headline" id=".E6.9C.8D.E5.8A.A1.E7.BA.A7.E9.94.99.E8.AF.AF.E4.BB.A3.E7.A0.81">服务级错误代码</span></h3>
<table border="1" cellspacing="0" cellpadding="0" width="100%" class="parameters" style="border-color: #CCCCCC;">
  <tbody><tr>
    <td width="10%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">错误代码</td>
    <td width="55%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">错误信息</td>
    <td width="35%" style="text-align:left;padding-left:5px;font-weight:bolder;border:1px solid #cccccc">详细描述</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20001</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IDs is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IDs参数为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20002</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Uid parameter is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Uid参数为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20003</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">User does not exists</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">用户不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20005</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Unsupported image type, only suport JPG, GIF, PNG</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不支持的图片类型，仅仅支持JPG、GIF、PNG</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20006</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Image size too large</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">图片太大</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20007</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Does multipart has image</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">请确保使用multpart上传图片</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20008</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Content is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">内容为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20009</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IDs is too many</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IDs参数太长了</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20012</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Text too long, please input text less than 140 characters</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">输入文字太长，请确认不超过140个字符</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20013</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Text too long, please input text less than 300 characters</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">输入文字太长，请确认不超过300个字符</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20014</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Param is error, please try again</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">安全检查参数有误，请再调用一次</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20015</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Account or ip or app is illgal, can not continue</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">账号、IP或应用非法，暂时无法完成此操作</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20016</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Out of limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">发布内容过于频繁</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20017</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Repeat content</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">提交相似的信息</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20018</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Contain illegal website</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">包含非法网址</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20019</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Repeat conetnt</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">提交相同的信息</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20020</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Contain advertising</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">包含广告信息</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20021</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Content is illegal</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">包含非法内容</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20022</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Your ip's behave in a comic boisterous or unruly manner</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">此IP地址上的行为异常</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20031</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Test and verify</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">需要验证码</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20032</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Update success, while server slow now, please wait 1-2 minutes</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">发布成功，目前服务器可能会有延迟，请耐心等待1-2分钟</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20101</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Target weibo does not exist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不存在的微博</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20102</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not your own weibo</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不是你发布的微博</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20103</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Can't repost yourself weibo</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不能转发自己的微博</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20104</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal weibo</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不合法的微博</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20109</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Weibo id is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">微博ID为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20111</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Repeated weibo text</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不能发布相同的微博</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20201</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Target weibo comment does not exist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不存在的微博评论</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20202</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal comment</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不合法的评论</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20203</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not your own comment</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不是你发布的评论</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20204</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Comment id is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">评论ID为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20301</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Can't send direct message to user who is not your follower</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不能给不是你粉丝的人发私信</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20302</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal direct message</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不合法的私信</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20303</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not your own direct message</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不是属于你的私信</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20305</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Direct message does not exist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不存在的私信</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20306</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Repeated direct message text</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不能发布相同的私信</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20307</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal direct message id</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">非法的私信ID</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20401</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Domain not exist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">域名不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20402</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Wrong verifier</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Verifier错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20501</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Source_user or target_user does not exists</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数source_user或者target_user的用户不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20502</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Please input right target user id or screen_name</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">必须输入目标用户id或者screen_name</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20503</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Need you follow user_id</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数user_id必须是你关注的用户</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20504</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Can not follow yourself </td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">你不能关注自己 </td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20505</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Social graph updates out of rate limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">加关注请求超过上限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20506</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Already followed</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">已经关注此用户</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20507</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Verification code is needed</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">需要输入验证码</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20508</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">According to user privacy settings,you can not do this</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">根据对方的设置，你不能进行此操作</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20509</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Private friend count is out of limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">悄悄关注个数到达上限 </td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20510</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not private friend</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不是悄悄关注人</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20511</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Already followed privately</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">已经悄悄关注此用户</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20512</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Please delete the user from you blacklist before you follow the user</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">你已经把此用户加入黑名单，加关注前请先解除</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20513</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Friend count is out of limit!</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">你的关注人数已达上限</td>
  </tr>
   <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20521</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Hi Superman, you have concerned a lot of people, have a think of how to make other people concern about you!&nbsp;! If you have any questions, please contact Sina customer service: 400 690 0000</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">hi 超人，你今天已经关注很多喽，接下来的时间想想如何让大家都来关注你吧！如有问题，请联系新浪客服：400 690 0000</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20522</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not followed </td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">还未关注此用户</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20523</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Not followers</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">还不是粉丝</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20524</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Hi Superman, you have cancelled concerning a lot of people, have a think of how to make other people concern about you!&nbsp;! If you have any questions, please contact Sina customer service: 400 690 0000</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">hi 超人，你今天已经取消关注很多喽，接下来的时间想想如何让大家都来关注你吧！如有问题，请联系新浪客服：400 690 0000</td>
  </tr><tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20601</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">List name too long, please input text less than 10 characters</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表名太长，请确保输入的文本不超过10个字符</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20602</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">List description too long, please input text less than 70 characters</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表描叙太长，请确保输入的文本不超过70个字符</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20603</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">List does not exists</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20604</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Only the owner has the authority</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不是列表的所属者</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20605</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal list name or list description</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表名或描叙不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20606</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Object already exists</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">记录已存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20607</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">DB error, please contact the administator</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">数据库错误，请联系系统管理员</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20608</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">List name duplicate</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表名冲突</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20610</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Does not support private list</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">目前不支持私有分组</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20611</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Create list error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">创建列表失败</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20612</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Only support private list</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">目前只支持私有分组</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20613</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">You hava subscriber too many lists</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">订阅列表达到上限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20614</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Too many lists, see doc for more info</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">创建列表达到上限，请参考API文档</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20615</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Too many members, see doc for more info</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">列表成员上限，请参考API文档</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20701</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Repeated tag text</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">不能提交相同的收藏标签</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20702</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Tags is too many</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">最多两个收藏标签</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20703</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Illegal tag name</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">收藏标签名不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20801</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Trend_name is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数trend_name是空值</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20802</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Trend_id is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数trend_id是空值</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20901</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Error: in blacklist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">错误:已经添加了黑名单</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20902</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Error: Blacklist limit has been reached.</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">错误:已达到黑名单上限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20903</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Error: System administrators can not be added to the blacklist.</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">错误:不能添加系统管理员为黑名单</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20904</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Error: Can not add yourself to the blacklist.</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">错误:不能添加自己为黑名单</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">20905</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Error: not in blacklist</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">错误:不在黑名单中</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21001</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Tags parameter is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">标签参数为空</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21002</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Tags name too long</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">标签名太长，请确保每个标签名不超过14个字符</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21101</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Domain parameter is error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数domain错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21102</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">The phone number has been used</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">该手机号已经被使用</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21103</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">The account has bean bind phone</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">该用户已经绑定手机</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21104</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Wrong verifier</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Verifier错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21301</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Auth faild</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">认证失败</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21302</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Username or password error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">用户名或密码不正确</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21303</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Username and pwd auth out of rate limit</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">用户名密码认证超过请求限制</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21304</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Version rejected</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">版本号错误</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21305</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Parameter absent</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">缺少必要的参数</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21306</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Parameter rejected</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">OAuth参数被拒绝</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21307</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Timestamp refused</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">时间戳不正确</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21308</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Nonce used</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数nonce已经被使用</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21309</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Signature method rejected</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">签名算法不支持</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21310</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Signature invalid</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">签名值不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21311</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Consumer key unknown</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数consumer_key不存在</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21312</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Consumer key refused</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数consumer_key不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21313</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Miss consumer key</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数consumer_key缺失</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21314</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token used</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token已经被使用</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21315</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token expired</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token已经过期</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21316</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token revoked</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21317</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token rejected</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Token不合法</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21318</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Verifier fail</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Pin码认证失败</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21319</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Accessor was revoked</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">授权关系已经被解除</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21320</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">OAuth2 must use https</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">使用OAuth2必须使用https</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21321</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Applications over the unaudited use restrictions</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">未审核的应用使用人数超过限制</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21327</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Expired token</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">token过期</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21335</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Request uid's value must be the current user</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">uid参数仅允许传入当前授权用户uid</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21501</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Urls is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数urls是空的</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21502</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Urls is too many</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数urls太多了</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21503</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">IP是空值</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21504</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Url is null</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">参数url是空值</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21601</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Manage notice error, need auth</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">需要系统管理员的权限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21602</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Contains forbid world</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">含有敏感词</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21603</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Applications send notice over the restrictions</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">通知发送达到限制</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21701</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Manage remind error, need auth</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">提醒失败，需要权限</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21702</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Invalid category</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">无效分类</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21703</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Invalid status</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">无效状态码</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">-</td>
  </tr>
  <tr>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">21901</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">Geo code input error</td>
    <td style="text-align:left;padding-left:5px;border:1px solid #cccccc">地理信息输入错误</td>
  </tr>
</tbody></table>
<p><br>
</p><p><br>
</p>
</div>