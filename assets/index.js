function loadTime () {
  var time = new Date();
  $('#date-time').text(`${time.getHours()}:${time.getMinutes()}`);
}

function sendRequest() {
 $('#send-request-btn').click(function() {
  document.cookie = "session" + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

  let webhook_url = $("#InputWebHookUrl").val();
  let username =    $("#InputUsername").val();
  let content =     $('#InputContent').val();
  let avatar_url =  $('#InputAvatarUrl').val();

  $.ajax({
    type: "POST",
    async: false,
    url: "/post-to-webhook",
    data: {
      webhook_url: webhook_url,
      username: username,
      content: content,
      avatar_url: avatar_url
    },
    success: function(result) {
      console.log("data, textStatus, XmlHttpRequest")
      $('#alert-danger').hide();
      $('#alert-success').show();
    },
    error: function(XMLHttpRequest, textStatus, errorThrown){
      $('#alert-success').hide();
      location.reload();
    },
    dataType: "json"
  });
 });
}

function loadSubmitBtn () {
  $('#send-request-btn').prop( "disabled", true );

  let form_webhook = $('#InputWebHookUrl');
  let form_username = $('#InputUsername');
  let form_content = $('#InputContent');
  let form_avatar_url = $('#InputAvatarUrl');

  let form_items = [form_webhook, form_username, form_content, form_avatar_url];
  form_items.forEach(element => {
    element.keyup(function(){
      if (form_webhook.val() != '' && form_username.val() != '' && form_content.val() != '' && form_avatar_url.val() != '') {
        $('#send-request-btn').prop( "disabled", false );
      } else {
        return;
      }
    });
  })
} 

function loadFormItemsIntoPreview() {
  let form_username = $('#InputUsername');
  let preview_username = $('#preview-username');
  form_username.keyup(function(){
    preview_username.text(form_username.val());
  })

  let form_content = $('#InputContent');
  let preview_content = $('#preview-content');
  form_content.keyup(function(){
    preview_content.text(form_content.val());
  })

  let form_avatar_url = $('#InputAvatarUrl');
  let preview_avatar_url = $('#preview-avatar_url');
  form_avatar_url.keyup(function(){
    preview_avatar_url.attr("src", form_avatar_url.val());
    
  })
}

function resetCookiesOnReload() {
  $(window).on('load', function(){
    document.cookie = "session" + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  });
}

resetCookiesOnReload();
loadTime();
loadSubmitBtn();
loadFormItemsIntoPreview();
sendRequest();
