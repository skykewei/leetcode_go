package com.leetcode.oj;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.Vector;

public class Intersect_350 {

  public int[] intersect(int[] nums1,int[] nums2){
    if(nums1.length==0||nums2.length==0){
      return new int[]{};
    }
    ArrayList<Integer> result = new ArrayList<>();
    HashMap<Integer, Integer> cntMap = new HashMap<>();
    for(int v:nums1){
      Integer mt = cntMap.get(v);
      if(mt!=null){
        cntMap.put(v,mt+1);
      }else{
        cntMap.put(v,1);
      }
    }
    for(int v2:nums2){
      Integer mt = cntMap.get(v2);
      if(mt!=null){
        if(mt>0){
          mt--;
          cntMap.put(v2,mt);
          result.add(v2);
        }
      }
    }
    int[] a = new int[result.size()];
    for(int i=0;i<a.length;i++){
      a[i] = result.get(i);
    }
    return a; 
  }
 
  public static void main(String[] args) {
    // TODO Auto-generated method stub

  }

}
