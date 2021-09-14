import express = require("express");
import { join } from "path";
const app = express();
const PORT = process.env.PORT || 8000;
app.use(express.urlencoded({ extended: false }));
app.use(express.json());
app.use(require("./routes/no-bitly"));
app.use(express.static(join(__dirname, "public")));
app.listen(PORT, () => console.log(`server on http://localhost:${PORT}`));
