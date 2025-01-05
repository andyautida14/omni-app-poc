document.body.addEventListener("htmx:afterRequest", function (evt) {
  const notifMsg = evt?.detail?.xhr?.getResponseHeader("X-Notif-Msg");
  if (notifMsg) {
    const status =
      evt.detail.xhr.getResponseHeader("X-Notif-Status") || "primary";
    UIkit.notification({
      message: notifMsg,
      pos: "bottom-right",
      status,
    });
  }
});

document.body.addEventListener("htmx:sendError", function () {
  UIkit.notification({
    message: "Server cannot be reached. Please try again later.",
    pos: "bottom-right",
    status: "danger",
  });
});
