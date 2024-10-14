import type { CustomThemeConfig } from "@skeletonlabs/tw-plugin";

export const theme: CustomThemeConfig = {
  name: "custom",
  properties: {
    // =~= Theme Properties =~=
    "--theme-font-family-base": `system-ui`,
    "--theme-font-family-heading": `Recoleta`,
    "--theme-font-color-base": "14 14 14",
    "--theme-font-color-dark": "255 255 255",
    "--theme-rounded-base": "9999px",
    "--theme-rounded-container": "24px",
    "--theme-border-base": "1px",
    // =~= Theme On-X Colors =~=
    "--on-primary": "240 240 240", // #f0f0f0
    "--on-secondary": "14 14 14", // #0e0e0e
    "--on-tertiary": "14 14 14", // #0e0e0e
    "--on-success": "14 14 14", // #0e0e0e
    "--on-warning": "14 14 14", // #0e0e0e
    "--on-error": "14 14 14", // #0e0e0e
    "--on-surface": "240 240 240", // #f0f0f0
    // =~= Theme Colors  =~=
    // primary | #4F774D
    "--color-primary-50": "229 235 228", // #e5ebe4
    "--color-primary-100": "220 228 219", // #dce4db
    "--color-primary-200": "211 221 211", // #d3ddd3
    "--color-primary-300": "185 201 184", // #b9c9b8
    "--color-primary-400": "132 160 130", // #84a082
    "--color-primary-500": "79 119 77", // #4F774D
    "--color-primary-600": "71 107 69", // #476b45
    "--color-primary-700": "59 89 58", // #3b593a
    "--color-primary-800": "47 71 46", // #2f472e
    "--color-primary-900": "39 58 38", // #273a26
    // secondary | #D89262
    "--color-secondary-50": "249 239 231", // #f9efe7
    "--color-secondary-100": "247 233 224", // #f7e9e0
    "--color-secondary-200": "245 228 216", // #f5e4d8
    "--color-secondary-300": "239 211 192", // #efd3c0
    "--color-secondary-400": "228 179 145", // #e4b391
    "--color-secondary-500": "216 146 98", // #D89262
    "--color-secondary-600": "194 131 88", // #c28358
    "--color-secondary-700": "162 110 74", // #a26e4a
    "--color-secondary-800": "130 88 59", // #82583b
    "--color-secondary-900": "106 72 48", // #6a4830
    // tertiary | #A4B494
    "--color-tertiary-50": "241 244 239", // #f1f4ef
    "--color-tertiary-100": "237 240 234", // #edf0ea
    "--color-tertiary-200": "232 236 228", // #e8ece4
    "--color-tertiary-300": "219 225 212", // #dbe1d4
    "--color-tertiary-400": "191 203 180", // #bfcbb4
    "--color-tertiary-500": "164 180 148", // #A4B494
    "--color-tertiary-600": "148 162 133", // #94a285
    "--color-tertiary-700": "123 135 111", // #7b876f
    "--color-tertiary-800": "98 108 89", // #626c59
    "--color-tertiary-900": "80 88 73", // #505849
    // success | #94C49F
    "--color-success-50": "239 246 241", // #eff6f1
    "--color-success-100": "234 243 236", // #eaf3ec
    "--color-success-200": "228 240 231", // #e4f0e7
    "--color-success-300": "212 231 217", // #d4e7d9
    "--color-success-400": "180 214 188", // #b4d6bc
    "--color-success-500": "148 196 159", // #94C49F
    "--color-success-600": "133 176 143", // #85b08f
    "--color-success-700": "111 147 119", // #6f9377
    "--color-success-800": "89 118 95", // #59765f
    "--color-success-900": "73 96 78", // #49604e
    // warning | #E7C586
    "--color-warning-50": "251 246 237", // #fbf6ed
    "--color-warning-100": "250 243 231", // #faf3e7
    "--color-warning-200": "249 241 225", // #f9f1e1
    "--color-warning-300": "245 232 207", // #f5e8cf
    "--color-warning-400": "238 214 170", // #eed6aa
    "--color-warning-500": "231 197 134", // #E7C586
    "--color-warning-600": "208 177 121", // #d0b179
    "--color-warning-700": "173 148 101", // #ad9465
    "--color-warning-800": "139 118 80", // #8b7650
    "--color-warning-900": "113 97 66", // #716142
    // error | #D36D56
    "--color-error-50": "248 233 230", // #f8e9e6
    "--color-error-100": "246 226 221", // #f6e2dd
    "--color-error-200": "244 219 213", // #f4dbd5
    "--color-error-300": "237 197 187", // #edc5bb
    "--color-error-400": "224 153 137", // #e09989
    "--color-error-500": "211 109 86", // #D36D56
    "--color-error-600": "190 98 77", // #be624d
    "--color-error-700": "158 82 65", // #9e5241
    "--color-error-800": "127 65 52", // #7f4134
    "--color-error-900": "103 53 42", // #67352a
    // surface | #4d4c47
    "--color-surface-50": "228 228 227", // #e4e4e3
    "--color-surface-100": "219 219 218", // #dbdbda
    "--color-surface-200": "211 210 209", // #d3d2d1
    "--color-surface-300": "184 183 181", // #b8b7b5
    "--color-surface-400": "130 130 126", // #82827e
    "--color-surface-500": "77 76 71", // #4d4c47
    "--color-surface-600": "69 68 64", // #454440
    "--color-surface-700": "58 57 53", // #3a3935
    "--color-surface-800": "46 46 43", // #2e2e2b
    "--color-surface-900": "38 37 35", // #262523
  },
};
