# Keoni's Biology Language (KBL)

Keoni's Biology Language is a language developed to effectively design, engineer, and build biological circuits. This is designed as the first iteration of a [Declarative Bioengineering Language](https://docs.google.com/document/d/1odf8q7ir9NsS0zPvArEg0j0WepEAOGAATOkvQpZgW7s/edit). The end goal of Keoni's Biology Language is to be able to define massively complicated virtual and physical operations on biological materials in a language that biologist can write.

Design goals:
- Extreme simplicity. A biologist with zero programming experience should be able to simulate expression of GFP in a new organism within 30 minutes of encountering the language.
- Useful. This language is not meant to do everything or represent everything. It is only scoped to accomplish specific experiments that are common in synthetic biology.
- Global Knowledge. Biology is wildly interconnected and there are massive databases serving useful data about biological devices. The KBL compiler should not only use local knowledge to make
 better decisions.

### What could you do with KBL?
For example, within a short file, a scientist should be able to define the GoldenGate modularization of an entire organism. An engineer should be able to ask for a single end product and a target organism, and the system will simulate different pathways with different enzymes, optimize and synthesis fix for each one of those enzymes, simulate the cloning into vectors for all those enzymes (including mutli-gene cassettes), while along the way simulating potential expression levels to gauge which RBSs and promoters should be used at each step of the process. 

Rather than a top-down approach to building this language where I try to think of every application it may be used in, I am taking a bottom-up approach, where I build the language according to my needs. As I build a toolkit for each new organism, I will incorporate features I need into KBL and into it's background library, Poly. Because it is specifically designed to fulfill my real needs, it is named "Keoni's Biology Language" rather than something more generic. 

Keoni
