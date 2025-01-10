// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileDeviceID The device ID.
type SyntheticsMobileDeviceID string

// List of SyntheticsMobileDeviceID.
const (
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD__2022__16_4                                          SyntheticsMobileDeviceID = "apple ipad (2022),16.4"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD__2022__17_3_1                                        SyntheticsMobileDeviceID = "apple ipad (2022),17.3.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_7TH_GEN__2019__13_3                                  SyntheticsMobileDeviceID = "apple ipad 7th gen (2019),13.3"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_9TH_GEN__2021__15_0_2                                SyntheticsMobileDeviceID = "apple ipad 9th gen (2021),15.0.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_9TH_GEN__2021__16_1                                  SyntheticsMobileDeviceID = "apple ipad 9th gen (2021),16.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_AIR__2022__15_4_1                                    SyntheticsMobileDeviceID = "apple ipad air (2022),15.4.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_MINI__5TH_GEN__14_6                                  SyntheticsMobileDeviceID = "apple ipad mini (5th gen),14.6"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_MINI__6TH_GEN__15_1                                  SyntheticsMobileDeviceID = "apple ipad mini (6th gen),15.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_11__2022__16_3                                   SyntheticsMobileDeviceID = "apple ipad pro 11 (2022),16.3"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2020__14_8                                 SyntheticsMobileDeviceID = "apple ipad pro 12.9 (2020),14.8"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2021__15_6_1                               SyntheticsMobileDeviceID = "apple ipad pro 12.9 (2021),15.6.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2022__16_5                                 SyntheticsMobileDeviceID = "apple ipad pro 12.9 (2022),16.5"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_MAX_13_1_3                                  SyntheticsMobileDeviceID = "apple iphone 11 pro max,13.1.3"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_13_6                                        SyntheticsMobileDeviceID = "apple iphone 11 pro,13.6"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_15_5                                        SyntheticsMobileDeviceID = "apple iphone 11 pro,15.5"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_13_3_1                                          SyntheticsMobileDeviceID = "apple iphone 11,13.3.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_14_0                                            SyntheticsMobileDeviceID = "apple iphone 11,14.0"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_16_3                                            SyntheticsMobileDeviceID = "apple iphone 11,16.3"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_MINI_14_2                                       SyntheticsMobileDeviceID = "apple iphone 12 mini,14.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_MINI_16_5                                       SyntheticsMobileDeviceID = "apple iphone 12 mini,16.5"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_MAX_14_5_1                                  SyntheticsMobileDeviceID = "apple iphone 12 pro max,14.5.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_14_5_1                                      SyntheticsMobileDeviceID = "apple iphone 12 pro,14.5.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_14_8                                        SyntheticsMobileDeviceID = "apple iphone 12 pro,14.8"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_15_0_2                                      SyntheticsMobileDeviceID = "apple iphone 12 pro,15.0.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_16_2                                        SyntheticsMobileDeviceID = "apple iphone 12 pro,16.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_2                                            SyntheticsMobileDeviceID = "apple iphone 12,14.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_6                                            SyntheticsMobileDeviceID = "apple iphone 12,14.6"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_8                                            SyntheticsMobileDeviceID = "apple iphone 12,14.8"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_15_6_1                                          SyntheticsMobileDeviceID = "apple iphone 12,15.6.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_16_6_1                                          SyntheticsMobileDeviceID = "apple iphone 12,16.6.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_MINI_15_0_2                                     SyntheticsMobileDeviceID = "apple iphone 13 mini,15.0.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_MINI_16_6                                       SyntheticsMobileDeviceID = "apple iphone 13 mini,16.6"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_MAX_15_1                                    SyntheticsMobileDeviceID = "apple iphone 13 pro max,15.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_MAX_17_3                                    SyntheticsMobileDeviceID = "apple iphone 13 pro max,17.3"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_15_0_2                                      SyntheticsMobileDeviceID = "apple iphone 13 pro,15.0.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_15_2                                        SyntheticsMobileDeviceID = "apple iphone 13 pro,15.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PLUS_16_1                                       SyntheticsMobileDeviceID = "apple iphone 14 plus,16.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_MAX_16_2                                    SyntheticsMobileDeviceID = "apple iphone 14 pro max,16.2"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_16_1                                        SyntheticsMobileDeviceID = "apple iphone 14 pro,16.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_17_3_1                                      SyntheticsMobileDeviceID = "apple iphone 14 pro,17.3.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_16_1                                            SyntheticsMobileDeviceID = "apple iphone 14,16.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_PRO_MAX_17_3_1                                  SyntheticsMobileDeviceID = "apple iphone 15 pro max,17.3.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_PRO_17_3_1                                      SyntheticsMobileDeviceID = "apple iphone 15 pro,17.3.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_17_2_1                                          SyntheticsMobileDeviceID = "apple iphone 15,17.2.1"
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_SE__2022__15_4_1                                   SyntheticsMobileDeviceID = "apple iphone se (2022),15.4.1"
	SYNTHETICSMOBILEDEVICEID_GALAXY_A40_9                                                    SyntheticsMobileDeviceID = "galaxy a40,9"
	SYNTHETICSMOBILEDEVICEID_GALAXY_A5_8_0_0                                                 SyntheticsMobileDeviceID = "galaxy a5,8.0.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_NOTE_10_9                                                SyntheticsMobileDeviceID = "galaxy note 10,9"
	SYNTHETICSMOBILEDEVICEID_GALAXY_NOTE5__AT_T__7_0                                         SyntheticsMobileDeviceID = "galaxy note5 (at&t),7.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S10_9                                                    SyntheticsMobileDeviceID = "galaxy s10,9"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S6_EDGE_SMNOT_G925F_7_0                                  SyntheticsMobileDeviceID = "galaxy s6 edge sm-g925f,7.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S8__TNOT_MOBILE__8_0_0                                   SyntheticsMobileDeviceID = "galaxy s8 (t-mobile),8.0.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S8_UNLOCKED_8_0_0                                        SyntheticsMobileDeviceID = "galaxy s8 unlocked,8.0.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9__UNLOCKED__9                                          SyntheticsMobileDeviceID = "galaxy s9 (unlocked),9"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9___UNLOCKED__8_0_0                                     SyntheticsMobileDeviceID = "galaxy s9+ (unlocked),8.0.0"
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9___UNLOCKED__9                                         SyntheticsMobileDeviceID = "galaxy s9+ (unlocked),9"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_2_8_0_0                                            SyntheticsMobileDeviceID = "google pixel 2,8.0.0"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_2_9                                                SyntheticsMobileDeviceID = "google pixel 2,9"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_XL_10                                            SyntheticsMobileDeviceID = "google pixel 3 xl,10"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_XL_9                                             SyntheticsMobileDeviceID = "google pixel 3 xl,9"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_10                                               SyntheticsMobileDeviceID = "google pixel 3,10"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_9                                                SyntheticsMobileDeviceID = "google pixel 3,9"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_XL_11                                           SyntheticsMobileDeviceID = "google pixel 3a xl,11"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_XL_12                                           SyntheticsMobileDeviceID = "google pixel 3a xl,12"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_10                                              SyntheticsMobileDeviceID = "google pixel 3a,10"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4__UNLOCKED__10                                    SyntheticsMobileDeviceID = "google pixel 4 (unlocked),10"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4__UNLOCKED__11                                    SyntheticsMobileDeviceID = "google pixel 4 (unlocked),11"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4_XL__UNLOCKED__10                                 SyntheticsMobileDeviceID = "google pixel 4 xl (unlocked),10"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4A_11                                              SyntheticsMobileDeviceID = "google pixel 4a,11"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4A_12                                              SyntheticsMobileDeviceID = "google pixel 4a,12"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_5A_5G_12                                           SyntheticsMobileDeviceID = "google pixel 5a 5g,12"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6_PRO_12                                           SyntheticsMobileDeviceID = "google pixel 6 pro,12"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6_12                                               SyntheticsMobileDeviceID = "google pixel 6,12"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6A_13                                              SyntheticsMobileDeviceID = "google pixel 6a,13"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_PRO_13                                           SyntheticsMobileDeviceID = "google pixel 7 pro,13"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_13                                               SyntheticsMobileDeviceID = "google pixel 7,13"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_14                                               SyntheticsMobileDeviceID = "google pixel 7,14"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_8_PRO_14                                           SyntheticsMobileDeviceID = "google pixel 8 pro,14"
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_8_14                                               SyntheticsMobileDeviceID = "google pixel 8,14"
	SYNTHETICSMOBILEDEVICEID_IPAD_8TH_GEN__2020__14_8                                        SyntheticsMobileDeviceID = "ipad 8th gen (2020),14.8"
	SYNTHETICSMOBILEDEVICEID_IPAD_AIR_2_13_6                                                 SyntheticsMobileDeviceID = "ipad air 2,13.6"
	SYNTHETICSMOBILEDEVICEID_IPAD_AIR_4TH_GEN__2020__14_8                                    SyntheticsMobileDeviceID = "ipad air 4th gen (2020),14.8"
	SYNTHETICSMOBILEDEVICEID_IPHONE_13_15_0_2                                                SyntheticsMobileDeviceID = "iphone 13,15.0.2"
	SYNTHETICSMOBILEDEVICEID_IPHONE_13_16_0_2                                                SyntheticsMobileDeviceID = "iphone 13,16.0.2"
	SYNTHETICSMOBILEDEVICEID_IPHONE_6_12_5_4                                                 SyntheticsMobileDeviceID = "iphone 6,12.5.4"
	SYNTHETICSMOBILEDEVICEID_IPHONE_6S_14_4_2                                                SyntheticsMobileDeviceID = "iphone 6s,14.4.2"
	SYNTHETICSMOBILEDEVICEID_IPHONE_7_14_8                                                   SyntheticsMobileDeviceID = "iphone 7,14.8"
	SYNTHETICSMOBILEDEVICEID_IPHONE_8_13_5_1                                                 SyntheticsMobileDeviceID = "iphone 8,13.5.1"
	SYNTHETICSMOBILEDEVICEID_IPHONE_8_14_4_2                                                 SyntheticsMobileDeviceID = "iphone 8,14.4.2"
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__13_6                                           SyntheticsMobileDeviceID = "iphone se (2020),13.6"
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__14_6                                           SyntheticsMobileDeviceID = "iphone se (2020),14.6"
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__15_0_2                                         SyntheticsMobileDeviceID = "iphone se (2020),15.0.2"
	SYNTHETICSMOBILEDEVICEID_IPHONE_X_14_6                                                   SyntheticsMobileDeviceID = "iphone x,14.6"
	SYNTHETICSMOBILEDEVICEID_IPHONE_XR_12_0                                                  SyntheticsMobileDeviceID = "iphone xr,12.0"
	SYNTHETICSMOBILEDEVICEID_IPHONE_XR_14_0                                                  SyntheticsMobileDeviceID = "iphone xr,14.0"
	SYNTHETICSMOBILEDEVICEID_IPHONE_XS_MAX_12_1                                              SyntheticsMobileDeviceID = "iphone xs max,12.1"
	SYNTHETICSMOBILEDEVICEID_IPHONE_XS_13_6                                                  SyntheticsMobileDeviceID = "iphone xs,13.6"
	SYNTHETICSMOBILEDEVICEID_LG_STYLO_5_9                                                    SyntheticsMobileDeviceID = "lg stylo 5,9"
	SYNTHETICSMOBILEDEVICEID_LG_STYLO_6_10                                                   SyntheticsMobileDeviceID = "lg stylo 6,10"
	SYNTHETICSMOBILEDEVICEID_MOTO_G_4_7_0                                                    SyntheticsMobileDeviceID = "moto g 4,7.0"
	SYNTHETICSMOBILEDEVICEID_MOTO_G7_PLAY_9                                                  SyntheticsMobileDeviceID = "moto g7 play,9"
	SYNTHETICSMOBILEDEVICEID_NEXUS_7_NOT__2ND_GEN__WIFI__6_0                                 SyntheticsMobileDeviceID = "nexus 7 - 2nd gen (wifi),6.0"
	SYNTHETICSMOBILEDEVICEID_ONEPLUS_8T_11                                                   SyntheticsMobileDeviceID = "oneplus 8t,11"
	SYNTHETICSMOBILEDEVICEID_PIXEL_2_XL_8_1_0                                                SyntheticsMobileDeviceID = "pixel 2 xl,8.1.0"
	SYNTHETICSMOBILEDEVICEID_PIXEL_2_XL_9                                                    SyntheticsMobileDeviceID = "pixel 2 xl,9"
	SYNTHETICSMOBILEDEVICEID_PIXEL_5_11                                                      SyntheticsMobileDeviceID = "pixel 5,11"
	SYNTHETICSMOBILEDEVICEID_PIXEL_5_12                                                      SyntheticsMobileDeviceID = "pixel 5,12"
	SYNTHETICSMOBILEDEVICEID_PIXEL_XL_8_0_0                                                  SyntheticsMobileDeviceID = "pixel xl,8.0.0"
	SYNTHETICSMOBILEDEVICEID_PIXEL_7_1_2                                                     SyntheticsMobileDeviceID = "pixel,7.1.2"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_A51_10                                                  SyntheticsMobileDeviceID = "samsung a51,10"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A10S_10                                          SyntheticsMobileDeviceID = "samsung galaxy a10s,10"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A13_5G_11                                        SyntheticsMobileDeviceID = "samsung galaxy a13 5g,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A53_5G_12                                        SyntheticsMobileDeviceID = "samsung galaxy a53 5g,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A7_8_0_0                                         SyntheticsMobileDeviceID = "samsung galaxy a7,8.0.0"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A71_11                                           SyntheticsMobileDeviceID = "samsung galaxy a71,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A73_5G_12                                        SyntheticsMobileDeviceID = "samsung galaxy a73 5g,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_J7__2018__8_0_0                                  SyntheticsMobileDeviceID = "samsung galaxy j7 (2018),8.0.0"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_NOTE20_11                                        SyntheticsMobileDeviceID = "samsung galaxy note20,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S20__UNLOCKED__10                                SyntheticsMobileDeviceID = "samsung galaxy s20 (unlocked),10"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S20___UNLOCKED__10                               SyntheticsMobileDeviceID = "samsung galaxy s20+ (unlocked),10"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_ULTRA_11                                     SyntheticsMobileDeviceID = "samsung galaxy s21 ultra,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_ULTRA_12                                     SyntheticsMobileDeviceID = "samsung galaxy s21 ultra,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21__11                                          SyntheticsMobileDeviceID = "samsung galaxy s21+,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_11                                           SyntheticsMobileDeviceID = "samsung galaxy s21,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_12                                           SyntheticsMobileDeviceID = "samsung galaxy s21,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_5G_12                                        SyntheticsMobileDeviceID = "samsung galaxy s22 5g,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_5G_13                                        SyntheticsMobileDeviceID = "samsung galaxy s22 5g,13"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_ULTRA_5G_12                                  SyntheticsMobileDeviceID = "samsung galaxy s22 ultra 5g,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22__5G_12                                       SyntheticsMobileDeviceID = "samsung galaxy s22+ 5g,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_ULTRA_13                                     SyntheticsMobileDeviceID = "samsung galaxy s23 ultra,13"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23__13                                          SyntheticsMobileDeviceID = "samsung galaxy s23+,13"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_13                                           SyntheticsMobileDeviceID = "samsung galaxy s23,13"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_14                                           SyntheticsMobileDeviceID = "samsung galaxy s23,14"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A_10_1_10                                    SyntheticsMobileDeviceID = "samsung galaxy tab a 10.1,10"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A_10_1_7_0                                   SyntheticsMobileDeviceID = "samsung galaxy tab a 10.1,7.0"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A8__2021__11                                 SyntheticsMobileDeviceID = "samsung galaxy tab a8 (2021),11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S4_8_1_0                                     SyntheticsMobileDeviceID = "samsung galaxy tab s4,8.1.0"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S6_9                                         SyntheticsMobileDeviceID = "samsung galaxy tab s6,9"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S7_11                                        SyntheticsMobileDeviceID = "samsung galaxy tab s7,11"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S8_12                                        SyntheticsMobileDeviceID = "samsung galaxy tab s8,12"
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_S20_ULTRA_10                                            SyntheticsMobileDeviceID = "samsung s20 ultra,10"
	SYNTHETICSMOBILEDEVICEID_SONY_XPERIA_XZ3_9                                               SyntheticsMobileDeviceID = "sony xperia xz3,9"
	SYNTHETICSMOBILEDEVICEID_XIAOMI_REDMI_NOTE_10_5G_11                                      SyntheticsMobileDeviceID = "xiaomi redmi note 10 5g,11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_10TH_GEN_2022_IOS_16        SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_10th_gen_2022_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_10TH_GEN_2022_IOS_17        SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_10th_gen_2022_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_9TH_GEN_2021_IOS_15         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_9th_gen_2021_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_AIR_2022_IOS_15             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_air_2022_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_MINI_5TH_GEN_IOS_14         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_mini_5th_gen_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_MINI_6TH_GEN_IOS_15         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_mini_6th_gen_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_11_2022_IOS_16          SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_pro_11_2022_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2020_IOS_14        SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_pro_12_9_2020_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2021_IOS_15        SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_pro_12_9_2021_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2022_IOS_16        SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_ipad_pro_12_9_2022_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_IOS_14                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_11_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_IOS_16                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_11_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_PRO_IOS_15             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_11_pro_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_14                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_15                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_16                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_MINI_IOS_14            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_mini_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_MINI_IOS_16            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_mini_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_PRO_IOS_14             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_pro_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_PRO_MAX_IOS_14         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_12_pro_max_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_MINI_IOS_15            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_13_mini_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_MINI_IOS_16            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_13_mini_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_IOS_15             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_13_pro_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_MAX_IOS_15         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_13_pro_max_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_MAX_IOS_17         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_13_pro_max_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_IOS_16                 SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_14_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PLUS_IOS_16            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_14_plus_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_IOS_16             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_14_pro_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_IOS_17             SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_14_pro_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_MAX_IOS_16         SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_14_pro_max_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_6_IOS_12                  SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_6_ios_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_8_IOS_13                  SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_8_ios_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_8_IOS_14                  SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_8_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_SE_2022_IOS_15            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_se_2022_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_SE_2022_IOS_16            SyntheticsMobileDeviceID = "synthetics:mobile:device:apple_iphone_se_2022_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GALAXY_A5_ANDROID_8                    SyntheticsMobileDeviceID = "synthetics:mobile:device:galaxy_a5_android_8"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GALAXY_NOTE5_ANDROID_7                 SyntheticsMobileDeviceID = "synthetics:mobile:device:galaxy_note5_android_7"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_3A_XL_ANDROID_11          SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_3a_xl_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4_UNLOCKED_ANDROID_11     SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_4_unlocked_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4_XL_UNLOCKED_ANDROID_10  SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_4_xl_unlocked_android_10"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4A_ANDROID_11             SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_4a_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6_ANDROID_12              SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_6_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6_PRO_ANDROID_12          SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_6_pro_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6A_ANDROID_13             SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_6a_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_ANDROID_13              SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_7_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_ANDROID_14              SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_7_android_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_PRO_ANDROID_13          SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_7_pro_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_8_ANDROID_14              SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_8_android_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_8_PRO_ANDROID_14          SyntheticsMobileDeviceID = "synthetics:mobile:device:google_pixel_8_pro_android_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPAD_AIR_2_IOS_13                      SyntheticsMobileDeviceID = "synthetics:mobile:device:ipad_air_2_ios_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPAD_AIR_4TH_GEN_2020_IOS_14           SyntheticsMobileDeviceID = "synthetics:mobile:device:ipad_air_4th_gen_2020_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_13_IOS_15                       SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_13_ios_15"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_13_IOS_16                       SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_13_ios_16"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_IOS_17                       SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_15_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_PRO_IOS_17                   SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_15_pro_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_PRO_MAX_IOS_17               SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_15_pro_max_ios_17"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_SE_2020_IOS_13                  SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_se_2020_ios_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_SE_2020_IOS_14                  SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_se_2020_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_X_IOS_14                        SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_x_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_XR_IOS_14                       SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_xr_ios_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_XS_IOS_13                       SyntheticsMobileDeviceID = "synthetics:mobile:device:iphone_xs_ios_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_LG_STYLO_6_ANDROID_10                  SyntheticsMobileDeviceID = "synthetics:mobile:device:lg_stylo_6_android_10"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_PIXEL_5_ANDROID_12                     SyntheticsMobileDeviceID = "synthetics:mobile:device:pixel_5_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_A51_ANDROID_10                 SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_a51_android_10"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_A71_ANDROID_11          SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_a71_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_NOTE20_ANDROID_11       SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_note20_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ANDROID_11          SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s21_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ANDROID_12          SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s21_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_PLUS_ANDROID_11     SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s21_plus_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ULTRA_ANDROID_11    SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s21_ultra_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_5G_ANDROID_12       SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s22_5g_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_5G_ANDROID_13       SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s22_5g_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_PLUS_5G_ANDROID_12  SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s22_plus_5g_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_ULTRA_5G_ANDROID_12 SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s22_ultra_5g_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ANDROID_13          SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s23_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ANDROID_14          SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s23_android_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_PLUS_ANDROID_13     SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s23_plus_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_PLUS_ANDROID_14     SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s23_plus_android_14"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ULTRA_ANDROID_13    SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_s23_ultra_android_13"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_A_10_1_ANDROID_7    SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_tab_a_10_1_android_7"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_S7_ANDROID_11       SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_tab_s7_android_11"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_S8_ANDROID_12       SyntheticsMobileDeviceID = "synthetics:mobile:device:samsung_galaxy_tab_s8_android_12"
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_XIAOMI_REDMI_NOTE_10_5G_ANDROID_11     SyntheticsMobileDeviceID = "synthetics:mobile:device:xiaomi_redmi_note_10_5g_android_11"
)

var allowedSyntheticsMobileDeviceIDEnumValues = []SyntheticsMobileDeviceID{
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD__2022__16_4,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD__2022__17_3_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_7TH_GEN__2019__13_3,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_9TH_GEN__2021__15_0_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_9TH_GEN__2021__16_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_AIR__2022__15_4_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_MINI__5TH_GEN__14_6,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_MINI__6TH_GEN__15_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_11__2022__16_3,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2020__14_8,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2021__15_6_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPAD_PRO_12_9__2022__16_5,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_MAX_13_1_3,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_13_6,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_PRO_15_5,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_13_3_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_14_0,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_11_16_3,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_MINI_14_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_MINI_16_5,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_MAX_14_5_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_14_5_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_14_8,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_15_0_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_PRO_16_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_6,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_14_8,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_15_6_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_12_16_6_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_MINI_15_0_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_MINI_16_6,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_MAX_15_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_MAX_17_3,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_15_0_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_13_PRO_15_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PLUS_16_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_MAX_16_2,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_16_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_PRO_17_3_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_14_16_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_PRO_MAX_17_3_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_PRO_17_3_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_15_17_2_1,
	SYNTHETICSMOBILEDEVICEID_APPLE_IPHONE_SE__2022__15_4_1,
	SYNTHETICSMOBILEDEVICEID_GALAXY_A40_9,
	SYNTHETICSMOBILEDEVICEID_GALAXY_A5_8_0_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_NOTE_10_9,
	SYNTHETICSMOBILEDEVICEID_GALAXY_NOTE5__AT_T__7_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S10_9,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S6_EDGE_SMNOT_G925F_7_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S8__TNOT_MOBILE__8_0_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S8_UNLOCKED_8_0_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9__UNLOCKED__9,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9___UNLOCKED__8_0_0,
	SYNTHETICSMOBILEDEVICEID_GALAXY_S9___UNLOCKED__9,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_2_8_0_0,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_2_9,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_XL_10,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_XL_9,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_10,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3_9,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_XL_11,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_XL_12,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_3A_10,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4__UNLOCKED__10,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4__UNLOCKED__11,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4_XL__UNLOCKED__10,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4A_11,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_4A_12,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_5A_5G_12,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6_PRO_12,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6_12,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_6A_13,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_PRO_13,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_13,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_7_14,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_8_PRO_14,
	SYNTHETICSMOBILEDEVICEID_GOOGLE_PIXEL_8_14,
	SYNTHETICSMOBILEDEVICEID_IPAD_8TH_GEN__2020__14_8,
	SYNTHETICSMOBILEDEVICEID_IPAD_AIR_2_13_6,
	SYNTHETICSMOBILEDEVICEID_IPAD_AIR_4TH_GEN__2020__14_8,
	SYNTHETICSMOBILEDEVICEID_IPHONE_13_15_0_2,
	SYNTHETICSMOBILEDEVICEID_IPHONE_13_16_0_2,
	SYNTHETICSMOBILEDEVICEID_IPHONE_6_12_5_4,
	SYNTHETICSMOBILEDEVICEID_IPHONE_6S_14_4_2,
	SYNTHETICSMOBILEDEVICEID_IPHONE_7_14_8,
	SYNTHETICSMOBILEDEVICEID_IPHONE_8_13_5_1,
	SYNTHETICSMOBILEDEVICEID_IPHONE_8_14_4_2,
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__13_6,
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__14_6,
	SYNTHETICSMOBILEDEVICEID_IPHONE_SE__2020__15_0_2,
	SYNTHETICSMOBILEDEVICEID_IPHONE_X_14_6,
	SYNTHETICSMOBILEDEVICEID_IPHONE_XR_12_0,
	SYNTHETICSMOBILEDEVICEID_IPHONE_XR_14_0,
	SYNTHETICSMOBILEDEVICEID_IPHONE_XS_MAX_12_1,
	SYNTHETICSMOBILEDEVICEID_IPHONE_XS_13_6,
	SYNTHETICSMOBILEDEVICEID_LG_STYLO_5_9,
	SYNTHETICSMOBILEDEVICEID_LG_STYLO_6_10,
	SYNTHETICSMOBILEDEVICEID_MOTO_G_4_7_0,
	SYNTHETICSMOBILEDEVICEID_MOTO_G7_PLAY_9,
	SYNTHETICSMOBILEDEVICEID_NEXUS_7_NOT__2ND_GEN__WIFI__6_0,
	SYNTHETICSMOBILEDEVICEID_ONEPLUS_8T_11,
	SYNTHETICSMOBILEDEVICEID_PIXEL_2_XL_8_1_0,
	SYNTHETICSMOBILEDEVICEID_PIXEL_2_XL_9,
	SYNTHETICSMOBILEDEVICEID_PIXEL_5_11,
	SYNTHETICSMOBILEDEVICEID_PIXEL_5_12,
	SYNTHETICSMOBILEDEVICEID_PIXEL_XL_8_0_0,
	SYNTHETICSMOBILEDEVICEID_PIXEL_7_1_2,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_A51_10,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A10S_10,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A13_5G_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A53_5G_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A7_8_0_0,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A71_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_A73_5G_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_J7__2018__8_0_0,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_NOTE20_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S20__UNLOCKED__10,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S20___UNLOCKED__10,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_ULTRA_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_ULTRA_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21__11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S21_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_5G_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_5G_13,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22_ULTRA_5G_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S22__5G_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_ULTRA_13,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23__13,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_13,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_S23_14,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A_10_1_10,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A_10_1_7_0,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_A8__2021__11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S4_8_1_0,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S6_9,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S7_11,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_GALAXY_TAB_S8_12,
	SYNTHETICSMOBILEDEVICEID_SAMSUNG_S20_ULTRA_10,
	SYNTHETICSMOBILEDEVICEID_SONY_XPERIA_XZ3_9,
	SYNTHETICSMOBILEDEVICEID_XIAOMI_REDMI_NOTE_10_5G_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_10TH_GEN_2022_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_10TH_GEN_2022_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_9TH_GEN_2021_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_AIR_2022_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_MINI_5TH_GEN_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_MINI_6TH_GEN_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_11_2022_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2020_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2021_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPAD_PRO_12_9_2022_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_11_PRO_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_MINI_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_MINI_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_PRO_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_12_PRO_MAX_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_MINI_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_MINI_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_MAX_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_13_PRO_MAX_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PLUS_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_14_PRO_MAX_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_6_IOS_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_8_IOS_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_8_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_SE_2022_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_APPLE_IPHONE_SE_2022_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GALAXY_A5_ANDROID_8,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GALAXY_NOTE5_ANDROID_7,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_3A_XL_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4_UNLOCKED_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4_XL_UNLOCKED_ANDROID_10,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_4A_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6_PRO_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_6A_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_ANDROID_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_7_PRO_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_8_ANDROID_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_GOOGLE_PIXEL_8_PRO_ANDROID_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPAD_AIR_2_IOS_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPAD_AIR_4TH_GEN_2020_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_13_IOS_15,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_13_IOS_16,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_PRO_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_15_PRO_MAX_IOS_17,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_SE_2020_IOS_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_SE_2020_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_X_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_XR_IOS_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_IPHONE_XS_IOS_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_LG_STYLO_6_ANDROID_10,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_PIXEL_5_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_A51_ANDROID_10,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_A71_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_NOTE20_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_PLUS_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S21_ULTRA_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_5G_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_5G_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_PLUS_5G_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S22_ULTRA_5G_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ANDROID_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_PLUS_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_PLUS_ANDROID_14,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_S23_ULTRA_ANDROID_13,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_A_10_1_ANDROID_7,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_S7_ANDROID_11,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_SAMSUNG_GALAXY_TAB_S8_ANDROID_12,
	SYNTHETICSMOBILEDEVICEID_SYNTHETICS_MOBILE_DEVICE_XIAOMI_REDMI_NOTE_10_5G_ANDROID_11,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileDeviceID) GetAllowedValues() []SyntheticsMobileDeviceID {
	return allowedSyntheticsMobileDeviceIDEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileDeviceID) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileDeviceID(value)
	return nil
}

// NewSyntheticsMobileDeviceIDFromValue returns a pointer to a valid SyntheticsMobileDeviceID
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileDeviceIDFromValue(v string) (*SyntheticsMobileDeviceID, error) {
	ev := SyntheticsMobileDeviceID(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileDeviceID: valid values are %v", v, allowedSyntheticsMobileDeviceIDEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileDeviceID) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileDeviceIDEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileDeviceID value.
func (v SyntheticsMobileDeviceID) Ptr() *SyntheticsMobileDeviceID {
	return &v
}
