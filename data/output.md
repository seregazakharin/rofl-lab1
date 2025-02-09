# База знаний 


## Лисов Алексей

**Вопрос:** Что представляет собой язык, распознаваемый недетерминированным конечным автоматом (НКА)?

**Ответ:** Язык, распознаваемый недетерминированным конечным автоматом (НКА) – это все такие слова,  по которым существует хотя бы один путь из стартовой вершины в терминальную.


**Вопрос:** Дан регулярный язык, опиши алгоритм нахождения кратчайшего слова, принадлежащего этому регулярному языку

**Ответ:** Регулярный язык может быть задан с помощью конечного автомата.  Так как автомат конечен, то мы можем его обойти (пройти через все состояния) за конечное время.  Так как нам нужно найти самое короткое слово, то эта задача сводится к тому, что нам необходимо  найти кратчайший путь от стартовой вершины до какой-либо терминальной. По определению это  можно сделать с помощью поиска в ширину (bfs, breadth-first search). Поиск в ширину - алгоритм, находящий  все кратчайшие пути от заданной вершины в невзвешенном графе. Запускаем поиск в ширину и выйдем из него,  когда пришли в терминальное состояние. Так как нам нужно явно найти кратчайшее слово, то после этого делаем  восстановление ответа. Это можно сделать используя дополнительную структуру данных (например массив prev), для того,  чтобы для каждого состояния v хранить состояние u, из которого мы в него пришли. Мы можем пройти от найденной  вершины по массиву prev, пока не придем в начальное состояние. Записав все переходы мы получим кратчайшее слово,  принадлежащее регулярному языку.


**Вопрос:** Опиши алгоритм подсчета количества слов определенной длины в заданном регулярном языке

**Ответ:** Обозначим регулярный язык за L и пусть длина слов, количество которых мы хотим найти - l. Так как язык L регулярен, то мы можем построить соответствующий ему конечный автомат A. Решим задачу с помощью динамического программирования. Пусть a_(q,i) – количество слов длины i, переводящих  автомат A из начального состояния q0 в состояние q.  Чтобы пересчитать эту величину, нужно просуммировать значения динамического программирования из предыдущего по длине слоя для всех состояний, из которых есть ребро в состояние q. Ответом является сумма элементов столбца, отвечающего за длину l, соответствующих терминальным вершинам.


**Вопрос:** Какими являются языки недетерминированных автоматов с магазинной памятью?

**Ответ:** Языки недетерминированных автоматов с магазинной памятью являются контекстно-свободными. То есть эти языки могут быть заданы с помощью контекстно свободных грамматик.


**Вопрос:** Какая структура данных может описать магазинную память у автомата с магазинной памятью?

**Ответ:** Магазинная память у автомата с магазинной памятью является стеком.

**Вопрос:** Опиши алгоритм нахождения эпсилон замыкания для каждой из вершин автомата?

**Ответ:** ε-замыкание состояния q – это множество состояний, достижимых из q только по ε-переходам.  Соотсветственно эпсилон замыкание для каждой из вершин автомата можно предподсчитать с помощью поиска в глубину  (dfs) для каждой вершины.


**Вопрос:** Дай определение произведения двух автоматов

**Ответ:** Прямым произведением двух ДКА A1=⟨Σ1,Q1,s1,T1,δ1⟩ и A2=⟨Σ2,Q2,s2,T2,δ2⟩ называется ДКА  A=⟨Σ,Q,s,T,δ⟩, где:  1) Σ = Σ1∪Σ2, то есть он работает над пересечением алфавитов двух данных автоматов  2) Q = Q1×Q2, множество пар состояниий включает в себя состояния обоих автоматов  3) s =⟨s1,s2), стартуем с символов в обоих автоматах  4) T=T1×T2, терминальные состояния включают в себя терминальные состояния обоих автоматов  5) δ(⟨q1,q2⟩,c)=⟨δ1(q1,c),δ2(q2,c)⟩, то есть переходим по символу в обоих автоматах"


**Вопрос:** Как из недетерминированного конечного автомата A сделать pushdown automat B?

**Ответ:** Для этого нужно заменить переход из состояния q в состояние p по символу x на такой же переход,  только добавить z_0/z_0, где z_0 это дно стека. Из этого следует, что регулярные языки являются подмножеством  МП-автоматных языков (языков автоматов с магазинной памятью).


**Вопрос:** Теорема Клини

**Ответ:** Теорема Клини гласит о том, что множество языков, принимаемых детерминированным конечным автоматом совпадает с  множеством языком, принимаемых академическим регулярным выражением.


**Вопрос:** Какой язык называется префиксным (беспрефиксным)

**Ответ:** Язык L называется префиксным, если для любого w не равного v из L не верно, что w – префикс v. Также такие языки называют беспрефиксными.


**Вопрос:** Что можно сказать о языке L, который принимается детерминированным конечным автоматом  с магазинной памятью по пустому стеку


**Ответ:** Это значит, что язык L принимается детерминированным конечным автоматом с магазинной памятью по терминальному  состоянию, а также язык L – префиксный


**Вопрос:** Является ли регулярным язык Дика с единственным типом скобок? (язык Дика - множество правильных скобочных структур вместе с пустой структурой,  образующее язык над алфавитом {a,b}.)


**Ответ:** Язык Дика не является регулярным. Докажем с помощью леммы о накачке. Предположим, что он регулярный, тогда по лемме о накачке существует  n, являющееся длиной накачки. Возьмём последовательность из n открывающих, а затем n закрывающих скобок.  Для неё существуют соответствующие x, y, z из леммы о накачке. Но так как |xy| <= n, то y  состоит только из открывающих скобок, причём по условию леммы y не пустая. А значит при i = 2 в строке xy^iz  получится больше открывающих скобок, чем закрывающих, то есть это будет не правильной скобочной последовательностью.  Получили противоречие. Следовательно язык Дика с единственным типом скобок не является регулярным.


**Вопрос:** Является ли контекстно-свободным языком разность контекстно-свободного и регулярного языка?


**Ответ:** Да, разность контекстно-свободного и регулярного языка является контекстно-свободным языком.



---

## Дмитрий Теньшов

**Вопрос:** Что такое переход машины Тьюринга?


**Ответ:** Переход машины Тьюринга — это функция, зависящая от состояния конечного управления и обозреваемого символа.  За один переход машина Тьюринга должна выполнить следующие действия: изменить состояние,  записать ленточный символ в обозреваемую клетку, сдвинуть головку влево или вправо.


**Вопрос:** Какие языки допускаются при помощи машины Тьюринга?


**Ответ:** Языки, допустимые с помощью машины Тьюринга, называются рекурсивно перечислимыми, или РП-языками.


**Вопрос:** Опишите прием «память в состоянии» машины Тьюринга.


**Ответ:** Память в состоянии — конечное управление можно использовать не только для представления позиции в «Программе» машины Тьюринга,  но и для хранения конечного объема данных.


**Вопрос:** Опишите прием «Подпрограммы» машины Тьюринга.


**Ответ:** Подпрограмма машины Тьюринга представляет собой множество состояний, выполняющее некоторый полезный процесс.  Это множество включает в себя стартовое состояние и еще одно состояние,  которое не имеет переходов и служит состоянием «возврата» для передачи управления какому-либо множеству состояний,  вызвавшему данную подпрограмму. «Вызов» подпрограммы возникает везде, где есть переход в ее начальное состояние.


**Вопрос:** Можно ли запомнить позицию ленточной головки в позиции управления у машины Тьюринга?


**Ответ:** Хотя позиции конечны в каждый момент времени, всё множество позиций может быть и бесконечным.  Если состояние должно представлять любую позицию головки, то в состоянии должен быть компонент данных,  имеющий любое целое в качестве значения. Из-за этого компонента множество состояний должно быть бесконечным,  даже если только конечное число состояний используется в любой конечный момент времени. Определение же машин Тьюринга требует,  чтобы множество состояний было конечным. Таким образом, запомнить позицию ленточной головки в конечном управлении нельзя.


**Вопрос:** Что такое счетчиковая машина?


**Ответ:** Счетчиковые машины — это класс машин, обладающий возможностью запоминать конечное число целых чисел (счетчиков)  и совершать различные переходы в зависимости от того, какие из счетчиков равны 0 (если таковые вообще есть).  Счетчиковая машина может только прибавить 1 к счетчику или вычесть 1 из него,  но отличить значения двух различных ненулевых счетчиков она не способна.


**Вопрос:** Какой язык допускается счетчиковой машиной?


**Ответ:** Каждый язык, допускаемый счетчиковой машиной, рекурсивно перечислим. Причина в том,  что счетчиковые машины являются частным случаем магазинных, а магазинные — частным случаем многоленточных машин Тьюринга,  которые по теореме допускают только рекурсивно перечислимые языки.


**Вопрос:** Допускается ли любой рекурсивно перечислимый язык двухсчетчиковой машиной?


**Ответ:** Для имитации машины Тьюринга и, следовательно, для допускания любого рекурсивно перечислимого языка, достаточно двух счетчиков.  Для обоснования этого утверждения вначале доказывается, что достаточно трех счетчиков,  а затем три счетчика имитируются с помощью двух.


**Вопрос:** Что такое универсальная машина Тьюринга?


**Ответ:** Универсальной машиной Тьюринга называют машину Тьюринга, которая может заменить собой любую машину Тьюринга.  Получив на вход программу и входные данные, она вычисляет ответ, который вычислила бы по входным данным машина Тьюринга,  чья программа была дана на вход.


**Вопрос:** Какое время необходимо многоленточной машине Тьюринга для имитации шагов компьютера?


**Ответ:** Рассмотрим компьютер, обладающий следующими свойствами: у него есть только инструкции,  увеличивающие максимальную длину слова не более чем на один; у него есть только инструкции,  которые многоленточная машина Тьюринга может выполнить на словах длиной k за O(k^2) или меньшее число шагов.  Шаг — это выполнение одной инструкции. Таким образом, выполнение n шагов работы компьютера можно проимитировать  на многоленточной машине Тьюринга с использованием не более O(n^3) шагов.


**Вопрос:** Как связаны мощности следующих машин Тьюринга: многодорожечная машина Тьюринга,  машина Тьюринга с односторонней лентой, многоленточная машина Тьюринга, недетерминированная машина Тьюринга?


**Ответ:** Многодорожечная машина Тьюринга, машина Тьюринга с односторонней лентой, многоленточная машина Тьюринга  и недетерминированная машина Тьюринга, несмотря на различия в их конструкции или правилах работы,  обладают одинаковой вычислительной мощностью, то есть способны вычислить одни и те же классы функций.  Различия между видами машин Тьюринга (например, между машинами с одним или несколькими лентами)  могут повлиять на эффективность вычислений (время или пространство), но не на саму вычислительную мощность.


**Вопрос:** Если проблема P1 неразрешима и ее можно свести к проблеме P2, то является ли проблема P2 неразрешимой?


**Ответ:** Если проблему P1 можно свести к проблеме P2 и если P1 неразрешима, то и P2 неразрешима.


**Вопрос:** Что такое рандомизированная машина Тьюринга?


**Ответ:** Рандомизированная машина Тьюринга — это вариант многоленточной машины Тьюринга. Первая лента,  как обычно для многоленточных машин, содержит вход. Вторая лента также начинается непустыми клетками.  В принципе, вся она содержит символы 0 и 1, выбранные с вероятностью 1/2. Вторая лента называется случайной лентой.  Третья и последующие, если используются, вначале пусты и при необходимости выступают как рабочие.



---

## Филимонов Михаил

**Вопрос:** Рекурсивные языки.

**Ответ:** Языки, допускаемые машинами Тьюринга, называются рекурсивно-перечислимыми (РП), а РП-языки, допускаемые МТ, которые всегда останавливаются, — рекурсивными. “Разрешимость” есть  синоним “рекурсивности”, однако языки чаще называются “рекурсивными”, а проблемы (которые  представляют собой языки, интерпретируемые как вопросы) — “разрешимыми”.  Если язык не является рекурсивным, то проблема, которую выражает этот язык, называется “неразрешимой”. Рекурсивный язык позволяет построить разрешающую функцию: т.е. МТ, возвращающую один из двух результатов (да-нет), и корректно завершающую работу.


**Вопрос:** Рекурсивно-перечислимые языки. Примеры языков, которые являются рекурсивно-перечислимыми, но не рекурсивными.

**Ответ:** Язык L является рекурсивно-перечислимым (РП-языком), если L = L(M) для некоторой машины Тьюринга M. Проблема останова машины Тьюринга является РП, но не рекурсивной. В действительности, определенная А. М. Тьюрингом  машина допускала, не попадая в допускающее состояние, а останавливаясь. Для МТ  M можно определить H(M) как множество входов w, на которых M останавливается  независимо от того, допускает ли M вход w. Тогда проблема останова состоит в опре  делении множества таких пар (M, w), у которых w принадлежит H(M). Это еще один  пример проблемы/языка, которая является РП, но не рекурсивной.


**Вопрос:** Что такое язык диагонализации L_d

**Ответ:** Язык диагонализации L_d — это множество всех цепочек w_i, не принадлежащих L(M_i). Понятие M_i, “i-й машины Тьюринга”. Это машина Тьюринга M, кодом которой является i-я двоичная цепочка w_i. В язык L_d  входит каждая цепочка в алфавите {0, 1}, которая, будучи проинтерпретированной как код МТ, не принадлежит языку этой МТ. Язык L_d является хорошим примером не РП-языка, т.е. его не допускает ни одна машина Тьюринга.


**Вопрос:** Является ли язык L_d рекурсивно-перечислимым

**Ответ:** Язык L_d не является рекурсивно-перечислимым, т.е. не существует машины Тьюринга, которая допускала бы L_d. Доказательство. Допустим, что L_d = L(M) для некоторой МТ M. Так как L_d — язык над алфавитом {0, 1}, M должна содержаться в построенной нами последовательности машин Тьюринга, поскольку эта последовательность содержит все МТ с входным алфавитом {0, 1}. Следовательно, в ней есть, по крайней мере, один код машины M, скажем, i, т.е. M = M_i. Понятие M_i, “i-й машины Тьюринга”. Это МТ M, кодом которой является i-я двоичная цепочка w_i. Выясним теперь, принадлежит ли w_i языку L_d. Если w_i принадлежит L_d, то M_i допускает w_i. Но тогда (по определению L_d) w_i не принадлежит L_d, так как L_d содержит лишь такие w_j, для которых M_j не допускает w_j. Точно так же, если w_i не принадлежит L_d, то M_i не допускает w_i. Но тогда (по определению L_d) w_i принадлежит L_d. Поскольку w_i не может одновременно и принадлежать, и не принадлежать L_d, приходим к противоречию с нашим предположением о том, что M существует. Таким образом, L_d не является рекурсивно-перечислимым языком.


**Вопрос:** Почему языки 'рекурсивные'?

**Ответ:** Рекурсивные функции были введены в 30-х годах XX века С. К. Клини.  Это название закрепилось за одним из наиболее распространённых вариантов  уточнения общего понятия арифметического алгоритма, то есть такого алгоритма,  допустимые исходные данные которого представляют собой системы натуральных чисел,  а возможные результаты применения — натуральные числа. Языки называются рекурсивными, если они являются рекурсивным подмножеством набора всех  возможных конечных последовательностей по алфавиту языка. Тезис Чёрча — Тьюринга: любая функция,  которая может быть вычислена физическим устройством, может быть вычислена машиной Тьюринга. То есть рекурсивные функции вычисляют те же классы алгоритмов, что и машины Тьюринга. Эквивалентно, формальный язык  является рекурсивным, если существует машина Тьюринга, которая при вводе конечной  последовательности символов всегда останавливается и принимает её, если она  принадлежит языку, и останавливается и отвергает её в противном случае.


**Вопрос:** Почему универсальный язык L_u является рекурсивно-перечислимым, но не рекурсивным.


**Ответ:** Язык L_u состоит из цепочек, интерпретируемых как код МТ, к которому дописан ее вход. Цепочка принадлежит L_u, если эта МТ до пускает данный вход. Язык L_u является рекурсивно-перечислимым. Допустим, что L_u рекурсивен. Тогда по теореме дополнение L_u (дополнение L_u)  также рекурсивный язык. Но если существует МТ M, допускающая L_u, то, используя описанный ниже метод,  можно построить МТ, допускающую L_d. Язык диагонализации L_d — это множество всех цепочек w_i, не принадлежащих L(M_i). Понятие M_i, “i-й машины Тьюринга”. Это МТ M, кодом которой является i-я двоичная цепочка w_i. В язык L_d  входит каждая цепочка в алфавите {0, 1}, которая, будучи проинтерпретированной как код МТ, не принадлежит языку этой МТ. Поскольку нам известно, что L_d не является РП, приходим к  противоречию с предположением, что язык L_u является рекурсивным. Предположим, что L(M) = L_u. Можно  преобразовать МТ M в МТ M', которая допускает L_d с помощью следующих действий: 1. M' преобразует  входную цепочку w в w111w. Легче это сделать, используя для копии w вторую ленту, и затем  преобразовать двухленточную МТ в одноленточную. 2. M' имитирует M на новом входе. Если w есть w_i  в нашем перечислении, то M' определяет, допускает ли M_i вход w_i. Поскольку M допускает L_u, то она  допускает тогда и только тогда, когда M_i не допускает w_i, т.е. когда w_i принадлежит L_d. Таким образом,  M' допускает w тогда и только тогда, когда w принадлежит L_d. Поскольку по теореме машины M' не существует,  приходим к выводу, что язык L_u не является рекурсивным.


**Вопрос:** Является ли язык, состоящий из кодов всех машин Тьюринга,  которые допускают хотя бы одну цепочку, L_ne рекурсивно-перечислимым/рекурсивным?


**Ответ:** L_ne – непустой язык состоит из кодов всех машин Тьюринга, которые допускают хотя бы одну цепочку.  Он рекурсивно-перечислим, но не рекурсивный. Чтобы доказать РП, достаточно предъявить МТ, допускающую L_ne.  Проще всего это сделать, описав недетерминированную МТ M. Работа M заключается в следующем: 1. На вход M подается  код МТ M_i. 2. Используя недетерминизм, M угадывает вход w, который, возможно, допускается M_i.  3. M проверяет, допускает ли M_i свой вход w. В этой части M может моделировать работу универсальной МТ U,  допускающей L_u (язык L_u состоит из цепочек, интерпретируемых как код МТ, к которому дописан ее вход. Цепочка принадлежит L_u, если эта МТ до пускает данный вход). 4. Если M_i допускает w, то и M допускает свой вход, т.е. M_i. Понятие M_i, “i-й машины  Тьюринга”. Это МТ M, кодом которой является i-я двоичная цепочка w_i. В этот язык входит каждая цепочка  в алфавите {0, 1}, которая, будучи проинтерпретированной как код МТ, не принадлежит языку этой МТ.  Таким образом, если M_i допускает хотя бы одну цепочку, то M угадает ее (среди прочих, конечно) и допустит M_i.  Если же L(M_i) = ∅, то ни одна из угаданных w не допускается M_i, и M не допустит M_i. Таким образом, L(M) = L_ne.


**Вопрос:** Является ли язык L_e, состоящий из кодов всех МТ, языки которых пусты, рекурсивно-перечислимым/рекурсивным?


**Ответ:** L_e – пустой язык состоит из кодов всех МТ, языки которых пусты. Он не рекурсивно-перечислим.


**Вопрос:** Теорема Райса


**Ответ:** Свойство называется тривиальным, если оно либо пустое (т.е. никакой язык вообще ему не удовлетворяет), либо содержит все РП-языки (рекурсивно-перечислимые). В противном случае свойство назы вается нетривиальным.  Теорема Райса: Всякое нетривиальное свойство языков, допускаемых МТ, является неразрешимым. Например, множество кодов  машин Тьюринга, допускающих пустой язык, согласно теореме Райса является неразрешимым. В действительности  этот язык не является РП, хотя его дополнение — множество кодов МТ, допускающих хотя бы одну цепочку, —  является РП, но не рекурсивным.


**Вопрос:** Существует ли такой вход для заданной машины Тьюринга,  при обработке которого машина выполняет более пяти переходов,  прежде чем достигнет состояния останова (либо завершится, либо зациклится)?


**Ответ:** Алгоритм решения становится очевидным, если заметить, что, когда МТ делает пять переходов,  она обозревает не более девяти клеток вокруг начальной позиции головки. Поэтому можно проимитировать  пять переходов МТ на любом из конечного числа входов, длина которых не более девяти. Если все эти  имитации не достигают останова, то делается вывод, что на любом входе данная МТ совершает более пяти переходов.


**Вопрос:** Что такое проблема соответствий Поста.


**Ответ:** Экземпляр проблемы соответствий Поста (ПСП) состоит из двух списков равной длины в некотором алфавите Σ.  Как правило, мы будем называть их списками A и B, и писать A = w_1, w_2, …, w_k и B = x_1, x_2, …, x_k при  некотором целом k. Для каждого i пара (w_i, x_i) называется парой соответствующих цепочек. Мы говорим, что  экземпляр ПСП имеет решение, если существует последовательность из одного или нескольких целых чисел  i_1, i_2, …, i_m, которая, если считать эти числа индексами цепочек и выбрать соответствующие цепочки из  списков A и B, дает одну и ту же цепочку, т.е. w_i_1w_i_2…w_i_m = x_i_1xi_2…x_i_m. В таком случае последовательность  i_1, i_2, …, i_m называется решающей последовательностью, или просто решением, данного экземпляра ПСП.


**Вопрос:** Что такое модифицированная проблема соответствий Поста.


**Ответ:** Модифицированной проблемой соответствий Поста, или МПСП.  В модифицированной ПСП на решение накладывается дополнительное требование,  чтобы первой парой в решении была пара первых элементов списков A и B.  Более формально, экземпляр МПСП состоит из двух списков A = w_1, w_2, …, w_k  и B = x_1, x_2, …, x_k, и решением является последовательность из 0 или  нескольких целых чисел i_1, i_2, …, i_m, при которой  w_1w_i_1w_i_2…w_i_m = x_1x_i_1x_i_2…x_i_m. Отметим, что цепочки обязательно начинаются парой (w_1, x_1),  хотя индекс 1 даже не указан в качестве начального элемента решения. Кроме того, в отличие от ПСП,  решение которой содержит хотя бы один элемент, решением МПСП может быть и пустая последовательность (когда w_1 = x_1).


**Вопрос:** Разрешима ли ПСП/проблема соответствий Поста


**Ответ:** Заданы два списка, содержащие одинаковое количество цепочек.  Спрашивается, можно ли, выбирая последовательности соответствующих  цепочек из этих двух списков, построить путем их конкатенации одну  и ту же цепочку. ПСП является важным примером неразрешимой проблемы.  Сводимость ПСП к ряду других проблем обеспечивает доказательство их неразрешимости


**Вопрос:** Разрешим ли вопрос о неоднозначности КС-грамматики (контекстно-свободной грамматики)


**Ответ:** Вопрос о неоднозначности КС-грамматики (контекстно-свободной грамматики) неразрешим.  Неразрешимые проблемы, связанные с контекстно-свободными языками.  Посредством сведения ПСП(проблема соответствий Поста) к вопросу об определении некоторых свойств КС-грамматики можно доказать неразрешимость многих вопросов  о КС-языках или их грамматиках. В силу неразрешимости ПСП это сведение доказывает неразрешимость про блемы неоднозначности КС-грамматики.  Так же, например, о включении одного КС-языка в  другой или о пустоте пересечения двух КС-языков."



---

## Пишикина Мария

**Вопрос:** Какие существуют приемы интерпретации ленты и конечного управления машины Тьюринга?


**Ответ:** Существует три приема интерпретации ленты и конечного управления машины Тьюринга: память в состоянии, многодорожечные ленты, подпрограммы.


**Вопрос:** Опишите прием «Многодорожечные ленты» в машине Тьюринга.


**Ответ:** При использовании приема «Многодорожечные ленты» рассматривается лента машины Тьюринга, образованная несколькими дорожками.  Каждая дорожка может хранить один символ (в одной клетке), и алфавит машины Тьюринга состоит из кортежей, с одним компонентом для каждой «дорожки».


**Вопрос:** Существуют ли не рекурсивно перечислимые языки, допускаемые многоленточными машинами Тьюринга?


**Ответ:** Каждый язык, допускаемый многоленточной машиной Тьюринга, рекурсивно перечислим.


**Вопрос:** Какое время необходимо одноленточной машине Тьюринга для имитации переходов многоленточной машины Тьюринга?


**Ответ:** Время, необходимое одноленточной машине Тьюринга для имитации n переходов многоленточной машины Тьюринга, есть O(n^2).


**Вопрос:** Что такое «мультистековая (многомагазинная) машина»?


**Ответ:** Мультистековая (многомагазинная) машина представляет собой детерминированный МП-автомат (или машина с магазинной памятью) с несколькими магазинами. Он получает свои входные данные, как и МП-автомат, из некоторого их источника,  а не с ленты или из магазина, как машина Тьюринга. Мультистековая машина имеет конечное управление, то есть конечное множество состояний, и конечный магазинный алфавит, используемый для всех магазинов.  Переход мультистековой машины основывается на состоянии, входном символе и верхних символах всех магазинов.


**Вопрос:** Как называется язык, который допускается односчетчиковой машиной?


**Ответ:** «Односчетчиковые машины» — это класс машин, которые могут запоминать значение одного целого числа («счетчика») и совершать различные переходы в зависимости от того, равен ли счетчик 0.  Односчетчиковая машина может только прибавить 1 к счетчику или вычесть 1 из него, но не способна различать различные ненулевые значения счетчика. Каждый язык, допускаемый односчетчиковой машиной, является КС-языком. Контекстно-свободный (КС) язык — это язык, задаваемый контекстно-свободной грамматикой.  Контекстно-свободной грамматикой называется грамматика, у которой в левых частях всех правил стоят только одиночные нетерминалы.


**Вопрос:** Можно ли имитировать машину Тьюринга на компьютере?

**Ответ:** Имитация машины Тьюринга на компьютере в принципе возможна, если допустить, что для имитации значащей части ленты существует потенциально бесконечный запас сменных запоминающих устройств вроде диска.  Поскольку физические ресурсы, необходимые для создания дисков, конечны, данный довод сомнителен. Однако, поскольку пределы памяти Вселенной неизвестны или, без сомнения, обширны,  предположение о бесконечности ресурсов (как для ленты машины Тьюринга) является практически реалистичным и в целом допустимо.


**Вопрос:** Какое время необходимо одноленточной машине Тьюринга для имитации переходов компьютера?


**Ответ:** Машина Тьюринга может имитировать n шагов компьютера за O(n^6) своих шагов.


**Вопрос:** Почему язык, допускающийся в недетерминированных машинах Тьюринга, также допускается и обычной детерминированной машиной Тьюринга?


**Ответ:** Несмотря на кажущуюся большую мощность недетерминированных машин в связи с тем, что они выполняют несколько возможных вычислений сразу  (требуя только, чтобы хоть одно из них заканчивалось в допускающем состоянии), любой язык, допускающийся недетерминированной машиной Тьюринга,  также допускается и обычной детерминированной машиной Тьюринга, поскольку она может моделировать любой недетерминированный переход, делая многократные копии состояния, если встречается неоднозначность.


**Вопрос:** Как машина Тьюринга имитирует реальный компьютер?


**Ответ:** Машина Тьюринга может имитировать память и управление реального компьютера путем использования одной ленты для записи всех элементов памяти и их содержимого — регистров, основной памяти, дисков и других запоминающих устройств.  Таким образом, можно быть уверенным, что все, не выполнимое машиной Тьюринга, не может быть сделано и компьютером.


**Вопрос:** Какая проблема может возникнуть при допуске языка в рандомизированную машину Тьюринга?


**Ответ:** Имея дело с рандомизированными машинами Тьюринга, нужно быть более аккуратным с тем, что значит допускание входа такой машиной; становится возможным, что машина Тьюринга не определяет функции математически корректно (т.е. независимо от рандомизированных данных).  Проблема в том, что при анализе действий рандомизированной машины Тьюринга со входом приходится рассматривать все возможные случайные последовательности на второй ленте. Вполне возможно, что машина Тьюринга  допускает при одних случайных последовательностях, но отвергает при других; в действительности, если рандомизированная машина Тьюринга должна делать что-то более эффективно, чем детерминированная машина Тьюринга,  то существенно, чтобы различные последовательности на рандомизированной ленте приводили к различному поведению.


**Вопрос:** Перечисли возможные операции над машинами Тьюринга.


**Ответ:** 1) Композиция (суперпозиция) машин Тьюринга - пусть две машины Тьюринга X и Y вычисляют функции f(P) и g(P) соответственно, тогда можно построить машину Тьюринга T = Y(X), вычисляющую суперпозицию функций g(f(P)).  2) Разветвление машин Тьюринга - пусть две машины Тьюринга X и Y вычисляют функции f(P) и g(P) соответственно, причём множеством значений функции f(P) является множество {0, 1}, тогда можно построить машину Тьюринга T = X->Y,  которая перерабатывает слово P в g(P), если f(P)=1 и оставляет его без изменений если f(P)=0.  3) Цикл машин Тьюринга - пусть две машины Тьюринга X и Y вычисляют функции f(P) и g(P) соответственно, причём множеством значений функции f(P) является множество {0, 1}, тогда можно построить машину Тьюринга T = X∘Y, которая  выполняет следующую последовательность действий: 
  1 - вычисляет f(P) и если f(P)=1, то вычисляет новое значение P=g(P), а если f(P)=0, то переходит к заключительному состоянию с выходным словом P; 
  2 - повторяет действие 1 до тех пор пока для очередного значения P не будет выполнено f(P)=0.


**Вопрос:** Что такое самоприменимая машина Тьюринга?


**Ответ:** Машина Тьюринга называется самоприменимой, если она останавливается, когда в качестве входного слова для неё используется описание самой машины.



---

