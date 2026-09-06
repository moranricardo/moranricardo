# PRISMA GLOBAL REPORT V2.1
| Repositorio | Estado | Accion Recomendada |
|---|---|---|
| moranricardo/moranricardo | OK P1 | P1 Seguridad refinada: sin termux en *.go,*.sh,*.py |
| moranricardo/moranricardo | BUILD FAIL | P2 Salud: corregir workflows con failure detectado |
| termux-app | GHOST DOCS | Solo .md, ignorado por filtro V2 (falso positivo eliminado) |
| moranricardo/cli | PENDIENTE SCAN | Validar en siguiente iteracion |

**Capas activas confirmadas:**
- P1 Seguridad: --include=*.go,*.sh,*.py,*.js,*.ts (ignora .md)
- P2 Salud: BUILD FAIL / OUTDATED / ALL CLEAR -> actual: BUILD FAIL
- P3 Estrategia: tabla estructurada + camo bust ?v=2
