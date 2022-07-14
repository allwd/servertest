export default {
  port: 3000,
  fetch(request) {
    if (request.headers.get("Authorization") !== "Test") {
      return new Response(JSON.stringify({ err: "Unauthorized." }));
    }
    return new Response(JSON.stringify({ status: "Success." }));
  },
};
