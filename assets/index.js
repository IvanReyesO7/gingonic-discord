console.log("Loaded")
$('#send-request-btn').click(function() {

  let webhook_url = $("#InputWebHookUrl").val();
  let username =    $("#InputUsername").val();
  let content =     $('#InputContent').val();
  let avatar_url =  $('#InputAvatarUrl').val();

  $.post( "/post-to-webhook",
  {
    webhook_url: webhook_url,
    username: username,
    content: content,
    avatar_url: avatar_url
  },
  "json"
  );
});