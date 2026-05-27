export default {
  async fetch(request, env) {
    return new Response(`[TNH_AI_PROJECT_01] Edge Service Active. DNA: ${env.PROJECT_ID}`);
  },
};
