import { z } from "zod";

const NodeVersion = z.union([
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		npm: z.string(),
		v8: z.string(),
		uv: z.string(),
		zlib: z.string(),
		openssl: z.string(),
		modules: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		npm: z.string(),
		v8: z.string(),
		uv: z.string(),
		zlib: z.string(),
		openssl: z.string(),
		modules: z.string(),
		lts: z.string(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		v8: z.string(),
		uv: z.string(),
		zlib: z.string(),
		openssl: z.string(),
		modules: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		v8: z.string(),
		uv: z.string(),
		openssl: z.string(),
		modules: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		v8: z.string(),
		uv: z.string(),
		modules: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		v8: z.string(),
		modules: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
	z.object({
		version: z.string(),
		date: z.string(),
		files: z.array(z.string()),
		v8: z.string(),
		lts: z.boolean(),
		security: z.boolean(),
	}),
]);
export type NodeVersion = z.infer<typeof NodeVersion>;

export const NodeVersions = z.array(NodeVersion);
export type NodeVersions = z.infer<typeof NodeVersions>;
