const { Router } = require("express");
const { exec } = require("child_process");
const router = Router();
const rateLimit = require("express-rate-limit");
const regex = /(http|https):\/\/(bit\.ly)([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?/;

const limiter = rateLimit({
  windowMs: 1 * 60 * 1000, // 1 minute
  max: 60, // limit 60 request for minute
  message: { out: "HTTP ERROR 429 to many requests" },
});

router.post("/no-bitly", limiter, (req: any, res: any) => {
  const { url } = req.body;
  if (url.match(regex)) {
    console.log(`input: ${url}`);
    exec(`URL=${url} ./no_bitly`, (stdout: String, stderr: String) => {
      res.json({ out: stderr });
    });
  } else {
    res.json({ out: `not a valid bitly url: ${url}` });
  }
});

module.exports = router;
