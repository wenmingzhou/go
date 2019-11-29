/**
 * 365运营后台通用页面JS
 */
$(function () {
  $("#form-clear").on('click',function () {
    window.location.href = window.location.href.split('?')[0];
  });
});