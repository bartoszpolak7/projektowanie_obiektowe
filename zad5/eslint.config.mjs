// domyślna konfiguracja ESLint dla TypeScript

import eslint from '@eslint/js'
import { defineConfig } from 'eslint/config'
import tseslint from 'typescript-eslint'

const eslintConfig = defineConfig([
    eslint.configs.recommended,
    tseslint.configs.recommended,
])

export default eslintConfig
