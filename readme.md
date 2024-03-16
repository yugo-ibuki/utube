
# 取得できる情報
| キー                                | サブキー                            | 説明                                         | 値           |
|-------------------------------------|-------------------------------------|---------------------------------------------|-------------|
| auditDetails                        |                                     | YouTubeパートナーの監査プロセスに関連するデータ  |             |
|                                     | communityGuidelinesGoodStanding     | コミュニティガイドラインの遵守状況           | True or False|
|                                     | contentIdClaimsGoodStanding         | 未解決の著作権クレームの有無                 | True or False|
|                                     | copyrightStrikesGoodStanding        | 著作権違反の有無                             | True or False|
| brandingSettings                    |                                     | チャンネルのブランディング情報                |             |
| brandingSettings.channel            |                                     | チャンネルビューのブランディングプロパティ    |             |
|                                     | country                             | チャンネルの国                               | A String    |
|                                     | defaultLanguage                     | チャンネルのデフォルト言語                   | A String    |
|                                     | defaultTab                          | ユーザーが表示するデフォルトのコンテンツタブ | A String    |
|                                     | description                         | チャンネルの説明                             | A String    |
|                                     | featuredChannelsTitle               | おすすめチャンネルタブのタイトル             | A String    |
|                                     | featuredChannelsUrls                | おすすめチャンネルURLリスト                  | [A String]  |
|                                     | keywords                            | チャンネルに関連するキーワード(カンマ区切り)  | A String    |
|                                     | moderateComments                    | ユーザー投稿コメントのモデレーション         | True or False|
|                                     | profileColor                        | チャンネルページの目立つ色                   | A String    |
|                                     | showBrowseView                      | 動画ブラウズタブの表示                       | True or False|
|                                     | showRelatedChannels                 | 関連チャンネルの提案                         | True or False|
|                                     | title                               | チャンネルタイトル                           | A String    |
|                                     | trackingAnalyticsAccountId          | チャンネルトラフィックを追跡するアナリティクスID | A String    |
|                                     | unsubscribedTrailer                 | 非購読者向けのトレーラー                     | A String    |
| brandingSettings.hints              |                                     | 実験的な追加ブランディングプロパティ          |             |
|                                     | property                            | プロパティ                                   | A String    |
|                                     | value                               | プロパティの値                               | A String    |
| brandingSettings.image              |                                     | チャンネルに関連するブランディング画像        |             |
| contentDetails                      |                                     | チャンネルのコンテンツに関する詳細情報        |             |
| contentDetails.relatedPlaylists     |                                     | 関連するプレイリスト情報                      |             |
|                                     | favorites                           | お気に入りのプレイリストID                    | A String    |
|                                     | likes                               | いいねした動画のプレイリストID                | A String    |
|                                     | uploads                             | アップロードした動画のプレイリストID          | A String    |
|                                     | watchHistory                        | 視聴履歴プレイリストID                        | A String    |
|                                     | watchLater                          | 後で見るプレイリストID                        | A String    |
| contentOwnerDetails               |                                   | YouTubeパートナーとリンクされたチャンネルの情報 |              |
|                                   | contentOwner                      | コンテンツオーナーのID                       | A String     |
|                                   | timeLinked                        | チャンネルがコンテンツオーナーにリンクされた日時 | A String     |
| conversionPings                   |                                   | チャンネルのコンバージョンピング情報         |              |
|                                   | pings                             | アプリが発火すべきピング情報                 | [詳細情報]   |
| etag                              |                                   | リソースのEtag                               | A String     |
| id                                |                                   | YouTubeがチャンネルを識別するためのID         | A String     |
| kind                              |                                   | リソースの種類                               | "youtube#channel" |
| localizations                     |                                   | 異なる言語のローカライゼーション情報         |              |
|                                   | a_key                             | チャンネルローカライゼーション設定           | [詳細情報]   |
| snippet                           |                                   | チャンネルの基本的な詳細情報                 |              |
|                                   | country                           | チャンネルの国                               | A String     |
|                                   | customUrl                         | チャンネルのカスタムURL                      | A String     |
|                                   | defaultLanguage                   | チャンネルのデフォルト言語                   | A String     |
|                                   | description                       | チャンネルの説明                             | A String     |
|                                   | localized                         | ローカライズされたタイトルと説明             | [詳細情報]   |
|                                   | publishedAt                       | チャンネル作成日時                           | A String     |
|                                   | thumbnails                        | チャンネルのサムネイル画像                   | [詳細情報]   |
|                                   | title                             | チャンネルのタイトル                         | A String     |
| statistics                        |                                   | チャンネルの統計情報                         |              |
|                                   | commentCount                      | チャンネルのコメント数                       | A String     |
|                                   | hiddenSubscriberCount             | 購読者数の表示/非表示                        | True or False|
|                                   | subscriberCount                   | チャンネルの購読者数                         | A String     |
|                                   | videoCount                        | アップロードされた動画の数                   | A String     |
|                                   | viewCount                         | チャンネルの視聴回数                         | A String     |
| status                            |                                   | チャンネルのステータス情報                   |              |
|                                   | isLinked                          | ユーザーがYouTubeユーザ名やG+アカウントにリンクしているか | True or False|
|                                   | longUploadsStatus                 | チャンネルの長時間アップロードステータス     | A String     |
|                                   | madeForKids                       | チャンネルが子供向けかどうか                 | True or False|
|                                   | privacyStatus                     | チャンネルのプライバシーステータス           | A String     |
|                                   | selfDeclaredMadeForKids           | 自己申告による子供向け設定                   | True or False|
| topicDetails                      |                                   | チャンネルに関連するFreebaseトピック情報     |              |
|                                   | topicCategories                   | チャンネルのコンテンツを説明するWikipedia URLのリスト | [A String] 