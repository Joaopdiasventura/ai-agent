# Confidence

[Índice](./README.md) | [English](../en/confidence.md)

`internal/confidence` calcula um score determinístico de auditoria. Não é uma probabilidade calibrada. O modo é `claim` quando o plano está ready e o reasoning é supported; é `abstention` quando o plano abstém ou o reasoning é insufficient.

O modo claim usa sinais ponderados: qualidade da query, concordância de retrieval, separação de ranking, força da evidência, directness da evidência, cobertura semântica, grounding do plano e grounding da resposta. Os pesos enfatizam força da evidência e cobertura semântica, ainda verificando que a resposta gerada foi autorizada pelo plano.

O modo abstention usa qualidade da query, ausência de evidência, grounding do plano e grounding da resposta. Ausência de evidência é mais forte quando o reasoning é insufficient e não há fatos relevantes fortes.

Qualidade da query combina idioma, intent, target, sinais semânticos e contagem de termos lexicais. Concordância de retrieval compara rankings de topo e overlap par a par. Separação mede a margem entre candidatos ranqueados ou grupos de reasoning. Grounding do plano verifica fatos planejados contra evidência de reasoning. Grounding da resposta verifica IDs de fatos gerados contra fatos planejados.

Scores são limitados a `[0,1]`. Níveis são high em `>=0.8`, medium em `>=0.6` e low abaixo disso.
