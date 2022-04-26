package nodeexpress

//TODO: MOVE TO GITHUB REPO/GISTS
const NodeExpressModelTemplate string = `const create{{ .DomainName }}Schema = {
    type: "object",
    properties: {},
  };
  const update{{ .DomainName }}Schema = {
    type: "object",
    properties: {},
  };
  
  module.exports = { create{{ .DomainName }}Schema, update{{ .DomainName }}Schema };`


const NodeExpressControllerTemplate string = `const { apiWrapper } = require("../../utils/https_connector");

async function get{{ .DomainName }}(body, pathParams, queryParams) {
  try {
    const options = {
      headers: {},
      url: "",
      data: {},
      method: "GET",
    };
    const response = await apiWrapper(options);
    return response;
  } catch (error) {
    console.log("Errored out in get{{ .DomainName }}", error);
    throw error;
  }
}

async function post{{ .DomainName }}(body, pathParams, queryParams) {
  try {
    const options = {
      headers: {},
      url: "",
      data: {},
      method: "GET",
    };
    const response = await apiWrapper(options);
    return response;
  } catch (error) {
    console.log("Errored out in get{{ .DomainName }}", error);
    throw error;
  }
}

async function update{{ .DomainName }}(body, pathParams, queryParams) {
  try {
    const options = {
      headers: {},
      url: "",
      data: {},
      method: "GET",
    };
    const response = await apiWrapper(options);
    return response;
  } catch (error) {
    console.log("Errored out in get{{ .DomainName }}", error);
    throw error;
  }
}

async function delete{{ .DomainName }}(body, pathParams, queryParams) {
  try {
    const options = {
      headers: {},
      url: "",
      data: {},
      method: "GET",
    };
    const response = await apiWrapper(options);
    return response;
  } catch (error) {
    console.log("Errored out in get{{ .DomainName }}", error);
    throw error;
  }
}

module.exports = {
  get{{ .DomainName }},
  update{{ .DomainName }},
  post{{ .DomainName }},
  delete{{ .DomainName }},
};
`

const NodeExpressRouteTemplate string = `const express = require("express");
const { validate } = require("express-jsonschema");

const router = express.Router();
const {
  get{{ .DomainName }},
  update{{ .DomainName }},
  post{{ .DomainName }},
  delete{{ .DomainName }},
} = require("../../controllers/{{ .DomainName | ToLower }}/{{ .DomainName | ToLower }}Controller");
const {
  create{{ .DomainName }}Schema,
  update{{ .DomainName }}Schema,
} = require("../../schemas/{{ .DomainName | ToLower }}/{{ .DomainName | ToLower }}Schema");
/* domain level level routing */
router.get("/{{ .DomainName | ToLower }}s", async (req, res, next) => {
  const { params, query, body } = req;
  const response = await get{{ .DomainName }}(body, params, query);
  res.json(response);
},);

router.post(
  "/{{ .DomainName | ToLower }}s",
  validate({ body: create{{ .DomainName }}Schema }),
  async (req, res, next) => {
    const { params, query, body } = req;
    const response = await post{{ .DomainName }}(body, params, query);
    res.json(response);
  },
);

router.put(
  "/{{ .DomainName | ToLower }}s/:id",
  validate({ body: update{{ .DomainName }}Schema }),
  async (req, res, next) => {
    const { params, query, body } = req;
    const response = await update{{ .DomainName }}(body, params, query);
    res.json(response);
  },
);

router.delete("/{{ .DomainName | ToLower }}s/:id", async (req, res, next) => {
  const { params, query, body } = req;
  const response = await delete{{ .DomainName }}(body, params, query);
  res.json(response);
},);

module.exports = router;
`
