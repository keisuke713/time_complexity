勉強会案

計算量の概念
->引数として与えられる入力値のサイズが増えるにつれ処理の時間がどれだけ増加するかを表した概念
->実際に何秒かかるかではなくどれくらいのペースで増加するかを表している
->例え達成したいことが同じでもコードによっては計算量が大きく変わることがある。

例)二部探索のやつをここで紹介

なぜ知っておいた方がいいのか
->計算量という概念、求め方を知っていた方が実装の時の判断基準が増える。特にデータが増えてくると自分が書いたコードがどのくらいの時間が処理に必要なのかなど知る必要が出てくると思う

よく見る計算量
・O(1)
->定数。
->インプットサイズがどれだけ増えようと処理時間が変わらない。そもそも引数を持たない。引数のデータ構造のサイズに依存しない処理。

・O(log n)
->対数時間というもの。インプットが増えても計算量は増加しにくい
->よく使われる例が二部探索、1から2^nの中から数字を思い浮かべてそれを当てるやつ
->1から順に聞いていくと最悪2^nかかるが二部探索ならn回の質問でいける
->数が大きくなるほど差は歴然で例えば1から1000の中で相手が思い浮かべた数を当てなければならない時線形探索だと最悪1000回質問しなければならないが、二分探索なら最大でも10回質問すれば答えを導き出せる(2 ^ 10 > 1000のため)

・O(n)
->リニア
->入力サイズが増加すると処理時間も線形で増加する。例 配列のfor文など

・O(n^2)
->ネストループなど。パッと思いつくのが九九の表の生成とか？？
->挿入ソートとか

求め方(係数と最大次数以外は省く
->どうしてそれはnが大きくなるにつれて些細な違いにしかならないから




よく見るデータ構造の計算量
(動的)配列(要素の挿入()、削除、探索、mapやfilter
->挿入 末尾はO(1)、配列がいっぱいで作り直すときは元の配列をコピーして作り直すからO(n)、配列の途中に入れる時はそれ以降の要素を全て後ろにずらす操作が必要になるからO(n)
->インデックスのアクセスはO(1)、要素をもとに配列に含まれているかなどするときは全探索が必要なのでO(n)
->削除は削除対象の要素より後ろの全ての要素を前にずらす操作が必要なのでO(n)
配列が得意としているのはインデックスによるアクセスで挿入や削除はあまり得意じゃないかも
インデックスアクセスをあまり必要とせず挿入や削除を多用するなら連結リストの方がいいかも
->mapやfilterといったメソッドも内部では要素を一つずつ見ているのO(n)の計算量

ハッシュ(要素の挿入、削除、探索、
基本的に全てO(1)?
なんで


全部が全部最適化すればいいというわけでもない多分
可読性とか諸々あるし

あとは二部探索
どこかでフィボナッチの例を入れよう


ちなみにhashmapはどうやってO(1)を実現しているの？
キーをハッシュ関数で変換して値を格納するところを決める実際の値は連結リストになっている
ただhashmapを