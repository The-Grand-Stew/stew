package serverless

const ServerlessNodeLambda string = `const log = require("lambda-log");
const handler = async (event, context, callback) => {
  try {
    log.options.meta.requestId = context.awsRequestId;
    log.info("{{ .FunctionName }} invoked");
    const response = {
      statusCode: 200,
      body: JSON.stringify({
        message: "{{ .FunctionName }} executed successfully!",
        input: event,
      }),
    };
    log.info("{{ .FunctionName }} execution complete");
    callback(null, response);
    // Use this code if you don't use the http event with the LAMBDA-PROXY integration
    // callback(null, { message: '{{ .FunctionName }} executed successfully!', event });
  } catch (error) {
    log.error("Error in {{ .FunctionName }} handler", error);
  }
};

module.exports = {handler};
`

const ServerlessNodeTest string = `'use strict';

// tests for {{ .FunctionName }}

const mod = require('./handlers/{{ .FunctionName }}/index');

const jestPlugin = require('serverless-jest-plugin');
const lambdaWrapper = jestPlugin.lambdaWrapper;
const wrapped = lambdaWrapper.wrap(mod, { handler: 'handler' });

describe('{{ .FunctionName }} sanity test', () => {
  beforeAll((done) => {
//  lambdaWrapper.init(liveFunction); // Run the deployed lambda

    done();
  });

  it('implement tests here', () => {
    return wrapped.run({}).then((response) => {
      expect(response).toBeDefined();
    });
  });
});
`
const ServerlessPackageJSON string = `
{
  "name": "{{ .Project }}",
  "version": "1.0.0",
  "description": "Lambda functions for {{ .Project }}",
  "main": "handler.js",
  "dependencies": {
    "lambda-log": "^3.1.0"
  }
}
`
