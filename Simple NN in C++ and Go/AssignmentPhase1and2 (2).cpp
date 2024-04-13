
#include <bits/stdc++.h>
/// THIS IS MY TEMPLATE ///
/*
Convert int to string

(1)

int a = 10;
char *intStr = itoa(a);
string str = string(intStr);

(2)

int a = 10;
stringstream ss;
ss << a;
string str = ss.str();

(3)

std::string s = std::to_string(42);
*/
using namespace std;



map < string, string > cgMap;





string paramdiff(string dx, map < int, pair <string, string> > cg)
{
    int len = cg.size();
    // cout<<"len : "<<len;
    int flag = 0;
    for(int i = 0; i<len; i++)
    {
        pair < string, string > p;
        p = cg[i];
        string str = p.second;
        //cout<<" str : " << str<<endl;
        int l = str.size();
        string ret = "";
        for(int j = 0; j<l; j++)
        {
            if(dx[0] == str[j])
            {
                // cout<<"dx[0] "<< dx[0]<<" str[j] "<< str[j] <<endl;
                flag = 1;
                if(j == 0) // in that expression the variable is first char
                {
                    if(str[j+1] == '*')
                    {
                        j+= 2;

                        while(j <l)
                        {
                            ret += str[j++];
                        }

                    }
                    else if(str[j+1] == '^')
                    {
                        if(str[j+2] > '1')
                        {
                            ret += str[j+2];
                            ret += '*';
                            ret += str[j];
                            ret += '^';
                            int numx = str[j+2] - '0';
                            numx--;
                            ret += to_string(numx);
                        }
                    }
                    else if(str[j+1] == '+')
                    {
                        ret = "1";
                    }
                    else if(str[j+1] == '-')
                    {

                        ret = "1";
                    }
                }
                return ret;
            }
        }
    }
    if (!flag)
    {
        return "didn't find";
    }
    return NULL;


}

//will calculate dy/dx


string differtiator2(int y, map < int, pair < string, string > > cg, int idx)
{
    string dy = "i" + to_string(y);//i8
    pair < string, string > p = cg[y];
    string term = p.second;
    int len = term.size();
    string newstr = "";
    string key = dy + "/";//i8/
    cout<<"mark1 ";
    cout<<term <<" "<<len<<" "<<dy<<" "<<key<<endl;



    for(int i = 0; i<len; i++)
    {
        string dx = "";

        int appOfI = 0, endOfI = 0;
        int flag = 0;
        if(term[i] == 'i')
        {
            flag = 1;
            appOfI = i;
            i++;
            dx += "i";
            while(term[i] >= '0' && term[i] <= '9')
            {
                dx += term[i];
                i++;
            }
            endOfI = i;
            i--;

            if(appOfI > 0)
            {
                if(term[appOfI-1] == '-')
                {
                    newstr += '-';
                }
                else if(term[appOfI-1] == '*')
                {
                    int it2 = appOfI-1;

                    while(it2 >= 0)
                    {
                        newstr += term[it2--];
                    }

                    reverse(newstr.begin(), newstr.end());

                }
            }

            if(term[endOfI] == '^')
            {
                int ii = endOfI + 1;
                int num = 0;
                while(term[ii] >= '0' && term[ii] <= '9')
                {
                    num *= 10;
                    num += term[ii] - '0';
                    ii++;
                }

                if(num == 2)
                {
                    newstr += "2*" + dx;
                }
            }
        }
        key += dx;

        if( newstr.size() == 1 && newstr[0] == '-' )
        {
            newstr += "1";
        }
        else if(!flag && newstr == "")
        {
            newstr += "1";
        }
        cgMap.insert(make_pair(key, newstr));


    }
    return "";
}
string differtiator(int y, int x, map < int, pair < string, string > > cg, int idx)
{
    string dy = "i" + to_string(y);
    string dx = "i" + to_string(x);
    string insStr = dy + "/" + dx;
    pair <string, string > p = cg[y];
    string term = p.second;

    int flag = 0;
    int pos = -1;
    int appi;


    int l = term.size();

    for(int i = 0; i<l; i++)
    {
        if(term[i] == 'i')
        {
            appi = i;
            string chek = "i";
            i++;

            while(term[i] >= '0' && term[i] <= '9')
            {
                chek += term[i];
                i++;
            }

            if(chek == dx)
            {
                flag = 1;
                pos = i;
            }
        }
    }
    string newstr;

    // cout<<"insStr : "<<insStr<<" flag "<<flag<<" pos "<<pos<<endl;
    int dum = pos;

    if(flag)
    {

        if( ((term[appi-1] == '+' || term[appi-1] == '-')  && pos == l) || l == 2)
        {
            // +i7 / -i7
            if(term[appi-1] == '-')
            {
                newstr += "-";
            }
            newstr += "1"; // 1 / -1
        }
        else if((term[pos] == '+' || term[pos] == '-'))
        {
            newstr = "1";

        }
        else if(term[pos] == '^')
        {
            int num = 0;
            pos++;
            while(pos < l)
            {
                num *= 10;
                char x = term[pos];
                num += x - '0';
                pos++;
            }

            newstr = to_string(num);

            newstr += "*";
            for(int i = 0; i<dum; i++)
            {
                newstr.push_back(term[i]);
            }

            if(num - 1 > 1)
            {
                newstr += "^";
                newstr += to_string(num-1);
            }
            // cout<<newstr<<endl;


        }
        else if(pos == l)
        {
            newstr = "";

            for(int i = 0; i<l; i++)
            {
                if(term[i] == '+' || term[i] == '-')
                {
                    newstr = "";
                }
                else if(term[i] == '*')
                {
                    break;
                }
                else
                {
                    newstr.push_back(term[i]);
                }
            }

            //  cout<<newstr<<endl;

        }

    }
    else
    {
        newstr = "0";
    }
    cgMap.insert(make_pair(insStr, newstr));
    return newstr;




}

double dp(string str, string str2, map <int, pair < string, string > > computationalGraph, map < string, double >   valuemap)
{
    cout<<endl;


    cout <<"str : "<<str ;
    cout <<" str2 : "<< str2 << endl;

    int num = 0;
    // cout<<"valuemap : "<<valuemap.size()<<endl;

    map < string, double > :: iterator itr;
    /*
        for (itr = valuemap.begin(); itr != valuemap.end(); itr++)
        {
            cout<<itr->first<<" "<<itr->second<<endl;
        }
    */

    int k = str.size();

    for (int i = 1 ; i<k; i++)
    {
        num *= 10;
        num += str[i] - '0';
    }
    int pop = num;
    //cout<<"num "<<num<<endl;

    pair < string, string > p;
    p = computationalGraph[num];

    string term = p.second;

    int l = term.size();

    // cout<<"first pass"<<endl;
// base case;
    for(int i = 0; i<l; i++)
    {
        if(term[i] == str2[0]) // have found the parameter
        {
            string ans = paramdiff(str2, computationalGraph );
            cout<<"paramdiff return string : "<<ans<<endl;

            if(ans[0] == 'i')
            {

                cout<<"valuemap[ans] : " <<valuemap[ans]<<endl;
                return valuemap[ans];


            }
            else
            {
                int ll = ans.size();
                int num2  = 0;
                for (int ij = 0; ij < ll ; ij++)
                {
                    if(ans[ij] >= '0' && ans[ij] <= '9')
                    {
                        num2 *= 10;
                        num2 += ans[ij] - '0';
                    }
                }

                if(ans[0] == '-')
                {
                    num2 *= -1;
                }
                cout<<"num2 : " <<endl;
                return num2;



            }

        }
    }

    //cout<<"second pass"<<endl;

    //haven't found the parameter

    l = term.size();

    cout<<"chk1 "<<term<<endl;
    double ans = 0;
    int flag = 0;
    int num1 = 1;

    for(int i = 0; i<l; i++)
    {
        if(term[i] == 'i') //
        {

            string st1 = "i";
            i++;
            while(term[i] >= '0' && term[i] <= '9')
            {
                st1 += term[i];
                i++;
            }
            i--;

            string st2 = "";
            st2 += str;
            st2 += '/';
            st2 += st1;

            cout<<"st2 : "<<st2<<endl;
            string diffterm = cgMap[st2];
            cout<<"diffterm : " <<diffterm<<endl;

            int len = diffterm.size();
            string st3 = "";

            for(int j = 0; j<len; j++)
            {

                if (diffterm[j] == 'i')
                {
                    flag = 1;
                    j++;
                    st3 = "i";
                    while(diffterm[j] >= '0' && diffterm[j] <= '9')
                    {
                        st3 += diffterm[j];
                        j++;
                    }
                    j--;

                }
            }

            if(diffterm.size() > 2)
            {
                num1 = diffterm[0] - '0';
                if(diffterm[1] >= '0' && diffterm[2] <= '1')
                {
                    num1 *= 10;
                    num1+= diffterm[1] - '0';
                }

            }


            if(flag)
            {
                ans = valuemap[st3];
                ans *= num1;
            }
            else
            {
                if(diffterm[0] == '1')
                {
                    ans = 1;
                }
                else if(diffterm[0] == '-')
                {
                    ans = -1;
                }
            }

        }

        if(ans == 0)
        {
            continue;
        }
        else
            break;
    }




    cout<<"ans : "<<ans<<endl;


    return ans * dp("i" + to_string(--pop), str2, computationalGraph, valuemap);




}


int main()
{
    freopen("input.txt", "r", stdin);
    //freopen("output.txt", "w", stdout);
    string str;
    string equation;
    int idx = 0;
    map < char, int > mymap;
    map < char, int > :: iterator itr;


    /*

    x      |        y
    2               4
    1               2
    3              -1
    4               5
    5              -2
    -1              3
    -2             -4
    6              11
    5              20
    */

    int xval = 2, wval = 1, bias = 3, aval = 2, bval = 3, cval = -1, yval = 4;
    int wvalarr[] = {1,-2,-1,1,-2,1,2,3,4,4};
    int bvalarr[] = {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1};
    int xarr [] = {2,1,3,4,5,-1,-2,6,5,-3};
    int yarr [] = {4,2,-1,5,-2,3,-4,11,20,9};

    mymap.insert(make_pair('^',5));
    mymap.insert(make_pair('/',4));
    mymap.insert(make_pair('*',3));
    mymap.insert(make_pair('+',2));
    mymap.insert(make_pair('-',1));

    while(cin>>equation)
    {
        cout<<equation<<endl;

        cout<< "x : "<<xval<<" y : "<<yval<<endl;
        cout<< "w : "<<wval<<" a : "<<aval<<endl;
        cout<< "b : "<<bval<<" c : "<<cval<<endl;

        //implementing shunting yard algorithm
        stack < char  > stk;
        queue < string > outputque;
        int len = equation.size();

        for( int i = 0; i<len; i++)
        {

            if((equation[i] >= 'a' && equation[i] <= 'z') ||(equation[i] >= '0' && equation[i] <= '9') )
            {

                if(equation[i] >= '0' && equation[i] <= '9' && equation[i+1] >= '0' && equation[i+1] <= '9')
                {
                    string str = "";
                    while(equation[i] >= '0' && equation[i] <= '9')
                    {
                        // cout<<"while 1"<<endl;
                        str.push_back(equation[i++]);
                    }
                    outputque.push(str);
                    i--;
                }
                else
                {
                    char sttt = equation[i];
                    string st4 = "";
                    st4.push_back(sttt);
                    outputque.push(st4);
                }
            }
            else if(equation[i] == '+' || equation[i] == '-' || equation[i] == '*' || equation[i] == '^' || equation[i] == '/' )
            {

                while(!stk.empty() && mymap[stk.top()] > mymap[equation[i]])
                {

                    // cout<<"while 2"<<endl;
                    string st = "";
                    st.push_back(stk.top());
                    stk.pop();
                    outputque.push(st);
                }
                stk.push(equation[i]);
            }
            else if(equation[i] == '(')
            {
                stk.push(equation[i]);
            }
            else if(equation[i] == ')')
            {
                while(stk.top() != '(')
                {

                    // cout<<"while 3"<<endl;
                    string st = "";
                    st.push_back(stk.top());
                    stk.pop();
                    outputque.push(st);

                }
                stk.pop();
            }

        }
        while(!stk.empty())
        {
            string st = "";
            st.push_back(stk.top());
            stk.pop();
            outputque.push(st);
        }

        cout<<"Queue :";
        len = outputque.size();
        while(len--)
        {
            //cout<< outputque.front()<<endl; outputque.pop();
            cout<<" -> "<<outputque.front();
            outputque.push(outputque.front());
            outputque.pop();
        }
        cout<<endl;


        //idea is this. I don't need to create node two store only 2 values. simply going to use map

        map < int, pair < string, string > > computationalGraph;
        map < int, pair < string, string > > :: iterator cgItr;

        //will create one element at a time
        int indices = 0;
        computationalGraph[indices++] = {"x", "i0"};

        //creating CG
        //steps ; - itr - insert queue - if op - pop queue - create string - insert back in queue - but put in map with key

        stack < string > secondstk;
        map < string, int > checkmap;

        map < string, int > :: iterator chktr;

        while(!outputque.empty())
        {
            string str = outputque.front();
            outputque.pop();

            //   cout<<"str : "<<str<<endl;
            char st = str[0]; // because str will be a single char
            itr = mymap.find(st); //checking if str is an op or not
            if(itr != mymap.end()) // if st == +/*/^
            {
                string fvar, svar, first;
                int idx2 = 0;
                if( !secondstk.empty())
                {
                    first = secondstk.top();
                    secondstk.pop();
                    fvar = first;
                    //   cout<<"fvar: "<< first <<endl;
                }
                if(!secondstk.empty())
                {
                    first = secondstk.top();
                    secondstk.pop();
                    svar = first;
                    //       cout<<"svar : "<<svar<<endl;
                }



                // the calculated 3x^2 is going back to queue
                int flag = 0;
                int cont = 0;

                pair < string, string > p;


                for(int ir = 0; ir< indices; ir++)
                {
                    p = computationalGraph[ir];



                    if(p.first == fvar)
                    {
                        //  cout<<"before swap "<<fvar<<" "<<svar<<endl;
                        if(fvar.size() == 1 && (fvar >= "0" && fvar <= "9") )
                            swap(fvar, svar);
                        // cout<<"after swat "<<fvar<<" "<<svar<<endl;
                        string dummy = "";
                        first = "";
                        first += svar;
                        first.push_back(st);
                        first += fvar;
                        flag = 1;
                        if(checkmap[svar] != 0)
                        {
                            dummy += "i";
                            dummy += to_string(checkmap[svar]);
                        }
                        else
                            dummy += svar;
                        dummy.push_back(st);
                        dummy += "i";
                        dummy += to_string(ir);


                        //   cout<<"pair first "<<p.first<<"|"<<"pair second "<<p.second<<endl;

                        //   cout<<"dummy " <<dummy<<" first "<<first<<endl;
                        secondstk.push(first);
                        checkmap[first] = indices;
                        computationalGraph[indices++] = {first, dummy};


                        //cout<<"cginsert 1 first , dummy " <<first<<" "<<dummy<<endl;
                        break;

                    }
                    else if( p.first == svar)
                    {
                        int flag2 = 0;
                        if(fvar.size() == 1 && (fvar >= "0" && fvar <= "9") )
                        {
                            swap(fvar, svar);
                            //  cout<<"swapdone"<<endl;
                            flag2 = 1;

                        }
                        string dummy = "";
                        first = "";
                        first += fvar;
                        first.push_back(st);
                        first += svar;
                        if(flag2)
                            swap(fvar, svar);
                        if(flag2)
                        {
                            flag = 1;
                            dummy += "i";
                            dummy += to_string(ir);
                            dummy.push_back(st);

                            if(checkmap[fvar] != 0)
                            {
                                dummy += "i";
                                dummy += to_string(checkmap[fvar]);
                            }
                            else
                                dummy += fvar;

                        }
                        else
                        {
                            flag = 1;
                            if(checkmap[fvar] != 0)
                            {
                                dummy += "i";
                                dummy += to_string(checkmap[fvar]);
                            }
                            else
                            {
                                dummy += fvar;

                            }
                            dummy.push_back(st);
                            dummy += "i";
                            dummy += to_string(ir);

                        }
                        checkmap[first] = indices;
                        // cout<<"dummy " <<dummy<<" first "<<first<<endl;
                        secondstk.push(first);
                        computationalGraph[indices++] = {first, dummy};
                        // cout<<"cginsert 2 first , dummy " <<first<<" "<<dummy<<endl;
                        break;
                    }
                }
                if(!flag)
                {
                    //  cout<<"if not: "<<first<<endl;
                    computationalGraph[indices++] = {first, "/"};
                }
            }
            else //secondque is for doing the operation on the variables
            {
                if((str >= "0" && str <= "9") && str.size() == 1)// for --3--
                {
                    string test = outputque.front();
                    char tt = test[0];
                    itr = mymap.find(tt);
                    if(itr != mymap.end() || secondstk.empty())//---3^ / ---3* / ---3+
                    {
                        if(test == "+" || test == "-" || test == "*" || test == "/"  )
                        {
                            string dummy = "";
                            dummy += secondstk.top();
                            secondstk.pop();
                            dummy += str;
                            secondstk.push(dummy);
                            // cout<< "rest of the cases secondstk push 0 "<<str <<endl;
                        }
                        else
                        {
                            secondstk.push(str);
                            //  cout<< "rest of the cases secondstk push 1 "<<str <<endl;
                        }

                    }
                    else
                    {
                        string dummy = "";
                        dummy = secondstk.top();
                        secondstk.pop();
                        dummy += str;
                        secondstk.push(dummy);
                        //  cout<< "rest of the cases secondstk push 2 "<<dummy <<endl;
                    }

                }
                else
                {
                    //   cout<< "rest of the cases secondstk push 4 "<<str <<endl;

                    secondstk.push(str);

                }
                //     cout<<"secondstk push : " << str <<" size : "<<secondstk.size()<<endl;

            }
        }
        computationalGraph[indices++] = { "y^", "i" + to_string(indices-1) };
        computationalGraph[indices++] = { "y-y^", "y-i" + to_string(indices-1)};
        computationalGraph[indices++] = { "(y-y^)^2", "i" + to_string(indices-1)+"^2"};

        stack < string >  checkstk = secondstk;
        pair < string, string > p;
        for ( cgItr = computationalGraph.begin(); cgItr != computationalGraph.end(); cgItr++)
        {
            p = cgItr->second;
            cout<<"i" <<cgItr->first <<": "<< p.first<<"              ||||||    "<<p.second<<" "<<endl;
        }


        cout<<endl<<endl;






        for (int i = indices - 1; i >= 0; i--)
        {
            string hulu = differtiator2(i, computationalGraph, indices);
        }


        cout<<endl<<endl;

        map <string, string > :: iterator cgMapItr;

        for(cgMapItr = cgMap.begin(); cgMapItr != cgMap.end(); cgMapItr++)
        {

            cout<<cgMapItr->first<<" "<<cgMapItr->second<<endl;
        }
        cout<<endl<<endl<<endl;

        map < string, double > valuemap;
        map < string, double >:: iterator itrval;
        valuemap.insert(make_pair("x", xval*1.0));
        valuemap.insert(make_pair("i0", valuemap["x"]));
        int cnt = 0;

        for ( int ij = 0; ij < indices; ij++)
        {
            pair < string, string > pairstr;
            pairstr = computationalGraph[ij];
            string pairstrstr = pairstr.second;

            int len = pairstrstr.size();

            double num1 = 0, num2 = 0;
            char opp = '/';

            int iflag = 0;
            int temp = 0;
            for(int i = 0; i<len; i++)
            {
                if(pairstrstr[i] == 'i')
                {
                    if(iflag == 1)
                    {
                        iflag = 2;
                    }
                    else
                        iflag = 1;

                    temp = num1;
                    string ss = "i";
                    i++;
                    while(pairstrstr[i] >= '0' && pairstrstr[i] <= '9')
                    {
                        ss.push_back(pairstrstr[i]);
                        i++;
                    }
                    // cout<<"ss : "<<ss<<" ";

                    if(valuemap.find(ss) != valuemap.end())
                    {
                        num1 = valuemap[ss];
                    }



                    i--;
                }

                if(pairstrstr[i] == 'w')
                {
                    num2 = wval;
                    //  cout<<"wval : " <<wval<<" ";
                }
                else if(pairstrstr[i] == 'a')
                {
                    num2 = aval;
                    //    cout<<"aval : " <<aval<<" ";
                }
                else if(pairstrstr[i] == 'b')
                {
                    num2 = bval;
                    //      cout<<"bval : " <<bval<<" ";

                }
                else if(pairstrstr[i] == 'c')
                {
                    num2 = cval;
                    //     cout<<"cval : " <<cval<<" ";

                }
                else if(pairstrstr[i] == 'y')
                {
                    num2 = yval;
                    //  cout<<"yval : " <<yval<<" ";
                }
                if(pairstrstr[i] == '+' || pairstrstr[i] == '*' || pairstrstr[i] == '^' || pairstrstr[i] == '-')
                {
                    if(pairstrstr[i] == '^')
                    {
                        opp = pairstrstr[i++];
                        num2 = pairstrstr[i];
                    }
                    else
                        opp = pairstrstr[i];
                }
            }


            double ans = 0;

            if(iflag == 2)
            {
                num2 = temp;
            }


            if(opp == '^')
            {
                ans = num1 * num1;
            }
            else
            {
                if(opp == '+')
                {
                    ans = num1 + num2;
                }
                else if(opp == '*')
                {
                    ans = num1 * num2;
                }
                else if( opp == '-')
                {
                    ans = num1 - num2;
                }
                else if(opp == '/')
                {
                    if(num1 != 0)
                    {
                        ans = num1;
                    }
                }
            }


            string inp = "i" + to_string(ij);

            valuemap.insert(make_pair(inp, ans));

            // cout<<inp<<" : num1 : "<<num1<<" num2 : "<<num2<<" opp : "<<opp<<endl;

        }
        cout<<"valuemap : "<<endl;

        for(itrval = valuemap.begin(); itrval != valuemap.end(); itrval++)
        {

            cout<<itrval->first<<"     "<<itrval->second<<endl;
        }
//// manually done
//        cout<<"diff wrt w : " << paramdiff("w", computationalGraph )<<endl;
//
//        cout<<"diff wrt a : " << paramdiff("a", computationalGraph )<<endl;
//
//        cout<<"diff wrt b : " << paramdiff("b", computationalGraph )<<endl;
//
//        cout<<"diff wrt c : " << paramdiff("c", computationalGraph )<<endl;
//
//// dp
//        cout<<indices<<endl;
//
//        double deltaparam1 = 0, deltaparam2 = 0;
//
//        str = "i" + to_string(indices-1);
//        string str2 = "w";
//
//       // cout<<"str ; "<<str<<endl<<endl;
//
//
//
//        if(paramdiff("w", computationalGraph ) != "didn't find")
//        {
//
//            deltaparam1 = dp(str, str2, computationalGraph,  valuemap);
//
//            cout<<"delta of w "<< deltaparam1<<endl;
//
//
//        }
//
//        if(paramdiff("a", computationalGraph ) != "didn't find")
//        {
//
//            deltaparam1 = dp(str, "a", computationalGraph,  valuemap);
//
//            cout<<"delta of a"<< deltaparam1<<endl;
//
//
//        }
//
//        if(paramdiff("b", computationalGraph ) != "didn't find")
//        {
//
//            deltaparam1 = dp(str, "b", computationalGraph,  valuemap);
//
//            cout<<"delta of b "<< deltaparam1<<endl;
//
//
//        }
//
//        if(paramdiff("c", computationalGraph ) != "didn't find")
//        {
//
//            deltaparam1 = dp(str, "c", computationalGraph,  valuemap);
//
//            cout<<"delta of c "<< deltaparam1<<endl;
//
//
//        }

        cgMap.clear();














        cout<<endl<<endl;
        cout<<"---------------------------------------------------------------------------"<<endl;
        cout<<"---------------------------------------------------------------------------"<<endl;
        cout<<"---------------------------------------------------------------------------"<<endl;
        cout<<endl<<endl;




    }
    return 0;
}
