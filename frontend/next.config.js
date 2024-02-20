loadEnv(process.env.NEXT_PUBLIC_APP_ENV);

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
};

module.exports = nextConfig;

function loadEnv(appEnv) {
  const defaultEnv = require(`./env/default.json`);
  const env = (() => {
    return appEnv !== undefined ? require(`./env/${appEnv}.json`) : {};
  })();

  Object.entries(defaultEnv).forEach(([key, value]) => {
    if (process.env[key] !== undefined) {
      return;
    }
    if (env[key] !== undefined) {
      value = env[key];
    }
    process.env[key] = value;
  });
}
