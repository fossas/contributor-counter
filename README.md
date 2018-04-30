# Bitbucket Contributor Counting

`contributors` counts active contributors in Bitbucket Server.

## Usage
`contributors` takes three parameters, all passed by environment variable:
- `BITBUCKET_URL`: the URL to your Bitbucket Server instance.
- `BITBUCKET_USERNAME`: your username on the Bitbucket Server instance.
- `BITBUCKET_PASSWORD`: your password on the Bitbucket Server instance.

```bash
BITBUCKET_URL=$YOUR_BITBUCKET_URL BITBUCKET_USER=$YOUR_BITBUCKET_USERNAME BITBUCKET_PASSWORD=$YOUR_BITBUCKET_PASSWORD contributors
```

`contributors` is tested on:
- Ubuntu 14.04 / Bitbucket 5.9.1 / git 2.17.0
- Ubuntu 14.04 / Bitbucket 5.9.1 / git 1.9.1

## Example
```bash
$ BITBUCKET_URL=http://172.17.0.1:7990 BITBUCKET_USER=leo BITBUCKET_PASSWORD=Y9iF6kxhNoBs87YNBms9 contributors
Found 112 contributors:
 187 jeffrey.mitchell@gmail.com
 182 kevin@fossa.io
  91 leo@fossa.io
  58 slackpad@gmail.com
  55 preetha@hashicorp.com
  44 jeff@hashicorp.com
  34 p.souchay@criteo.com
  33 leo@leozhang.me
  32 banks@banksco.de
  30 kylehav@gmail.com
  25 briankassouf@users.noreply.github.com
  25 yoko@hashicorp.com
  18 jackpearkes@gmail.com
  17 mjkeeler7@gmail.com
  17 vishalnayak@users.noreply.github.com
  15 cleung2010@gmail.com
  14 mkeeler@users.noreply.github.com
  11 vishalnayakv@gmail.com
  10 anuccio1@gmail.com
   9 jathomps08@gmail.com
   7 christopher.hoffman@gmail.com
   7 devin@twuni.org
   6 andy@hashicorp.com
   6 brianshumate@users.noreply.github.com
   6 edward.steel@gmail.com
   4 brian.nuszkowski@me.com
   4 jim@kalafut.net
   4 seth@sethvargo.com
   3 gh.je@mailhero.io
   3 giaquinti@slack-corp.com
   3 kostozyb@gmail.com
   3 meirish@users.noreply.github.com
   3 rberlind@optonline.net
   2 32205350+jeis2497052@users.noreply.github.com
   2 7660718+yhyakuna@users.noreply.github.com
   2 alex.dadgar@gmail.com
   2 burdandrei@users.noreply.github.com
   2 gobins@users.noreply.github.com
   2 jgiles@paxos.com
   2 robison@users.noreply.github.com
   2 seemethere101@gmail.com
   2 xiangli.cs@gmail.com
   1 34144035+seanjfellows@users.noreply.github.com
   1 5408930+rin1221@users.noreply.github.com
   1 Harrisonbro@users.noreply.github.com
   1 a.g.aleksandrov@gmail.com
   1 alex@hashicorp.com
   1 alexandrumd@yahoo.com
   1 alvin@hashicorp.com
   1 anubhavmishra@me.com
   1 arthur.lutz@logilab.fr
   1 avoidik@gmail.com
   1 bartlettc@gmail.com
   1 bharathgowda23@gmail.com
   1 bonifaido@gmail.com
   1 brian@brianshumate.com
   1 burdandrei@gmail.com
   1 choffman@hashicorp.com
   1 dev@marienfressinaud.fr
   1 dominikmueller@users.noreply.github.com
   1 emilyye@google.com
   1 ezyang@mit.edu
   1 gechr@users.noreply.github.com
   1 georgec.perez@gmail.com
   1 github@nafn.de
   1 guillaume@paralint.com
   1 hana378@gmail.com
   1 hannah@hashicorp.com
   1 huang.s.alvin@gmail.com
   1 i@cikenerd.com
   1 ja@tbedrich.cz
   1 jagiello.lukasz@gmail.com
   1 jcowen@hashicorp.com
   1 jcrowthe@users.noreply.github.com
   1 jed.s.bradshaw@gmail.com
   1 jeff@immutability.io
   1 jescalan@users.noreply.github.com
   1 jeweljar@hanmail.net
   1 jkalafut@hashicorp.com
   1 johncowen@users.noreply.github.com
   1 jscaltreto@users.noreply.github.com
   1 jsoref@users.noreply.github.com
   1 kevin@kevinwang.com
   1 kieran.othen@mac.com
   1 kpaulisse@users.noreply.github.com
   1 maxiwalther@mac.com
   1 mcintoshj@gmail.com
   1 megbeguk@gmail.com
   1 mitchell.hashimoto@gmail.com
   1 mohsen0@users.noreply.github.com
   1 ntroncos@gmail.com
   1 paddy@carvers.co
   1 platt.nicholas@gmail.com
   1 ppawelslomka@gmail.com
   1 przemyslaw.dabek@gmail.com
   1 public@paulstack.co.uk
   1 rmbrad@gmail.com
   1 robert.kreuzer@gmail.com
   1 runsisi@zte.com.cn
   1 samnap+github@gmail.com
   1 shihovmy@gmail.com
   1 slopeinsb@users.noreply.github.com
   1 sryabkov@users.noreply.github.com
   1 sschuberth@users.noreply.github.com
   1 thibaut.rousseau@protonmail.com
   1 tomwilkie@users.noreply.github.com
   1 trott@odaacabeef.com
   1 v-karbovnichy@users.noreply.github.com
   1 wim@42.be
   1 xmitchx@gmail.com
   1 y.fouquet@criteo.com
   1 yfouquet@localhost.localdomain

```