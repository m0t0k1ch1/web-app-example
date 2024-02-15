declare namespace NodeJS {
  interface ProcessEnv extends Readonly<typeof import("./default.json")> {}
}
